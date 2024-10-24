package docs

import (
	dto "docs/internal/Dto"
	grpc_client "docs/internal/grpc-client"
	"docs/internal/models"
	"docs/internal/services"
	"log"

	"github.com/google/uuid"
)

func CreateDoc(docPost dto.DocPost, res chan<- interface{}) {
	result := make(chan models.ResultChan[uuid.UUID], 1)
	committed := make(chan bool, 2) // Buffered channel to prevent goroutine leaks
	insert := services.Service.Conne.DbInsert

	docIDChan := grpc_client.GrpcClient(docPost)
	docID := <-docIDChan
	newDoc := models.Document{
		DocumentName: docPost.DocName,
		MongoID:      docID,
	}

	go func() {
		newDoc.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting document: %v", insertResult.Error)
		res <- uuid.Nil
		close(res)
		return
	}

	document_contribution := models.DocumentContribution{
		UserID:     docPost.UserUuid,
		DocumentID: insertResult.Data,
		Role:       "editor",
	}

	go CreateDocContribution(document_contribution, committed)
	go CreateDocOwner(docPost.UserUuid, insertResult.Data, committed)

	success := true
	for i := 0; i < 2; i++ {
		if !<-committed {
			success = false
		}
	}
	close(committed) 

	if !success {
		res <- uuid.Nil
	} else {
		res <- insertResult.Data
	}
	close(res)
}
