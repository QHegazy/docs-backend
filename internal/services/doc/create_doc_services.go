package docs

import (
	dto "docs/internal/Dto"
	grpc_client "docs/internal/grpc-client"
	"docs/internal/models"
	"docs/internal/services"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type UserDoc struct {
	UserID     uuid.UUID
	DocumentID uuid.UUID
}

func CreateDoc(docPost dto.DocPost, public dto.Visibility, res chan<- interface{}) {
	result := make(chan models.ResultChan[uuid.UUID], 1)
	committed := make(chan bool, 4)
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
	switch public {
	case 1:
		view := models.DocumentPublicView{
			DocumentID: insertResult.Data,
		}
		go CreateDocPublicView(view, committed)

	case 2:
		edit := models.DocumentPublicEdit{
			DocumentID: insertResult.Data,
		}
		go CreateDocPublicEdit(edit, committed)

	}
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

func QueryDoc(userDoc UserDoc, res chan<- *models.Document) {
	query := services.Service.Conne.DbRead
	newDoc := models.Document{
		DocumentID: userDoc.DocumentID,
	}

	result := make(chan models.ResultChan[*models.Document])

	go func() {
		newDoc.Query(query, result)
		close(result)
	}()

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying document: %v", queryResult.Error)
		res <- nil
		return
	}
	res <- queryResult.Data
	defer close(res)
	defer close(result)

}
func DeleteDoc(userDoc UserDoc, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	check := make(chan bool, 1)
	result := make(chan models.ResultChan[error])
	newDoc := models.Document{
		DocumentID: userDoc.DocumentID,
	}
	go CheckUserOwnDoc(userDoc, check)

	go func() {
		if <-check {
			fmt.Println("Check:", check)
			newDoc.Delete(delete, result)
			close(result)
		}
	}()

	deleteResult := <-result
	if deleteResult.Error != nil {
		log.Printf("Error deleting document: %v", deleteResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

func CheckUserOwnDoc(userDoc UserDoc, res chan<- bool) {
	query := services.Service.Conne.DbRead
	newDoc := models.DocumentOwnership{
		UserID:     userDoc.UserID,
		DocumentID: userDoc.DocumentID,
	}

	result := make(chan models.ResultChan[*models.DocumentOwnership])

	go func() {
		newDoc.Query(query, result)
		close(result)
	}()

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying document ownership: %v", queryResult.Error)
		res <- false
		return
	}
	res <- queryResult.Data != nil
	defer close(res)
	defer close(result)
}
