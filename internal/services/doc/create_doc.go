package docs

import (
	dto "docs/internal/Dto"
	"docs/internal/models"
	"docs/internal/services"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func CreateDoc(docPost dto.DocPost, res chan<- interface{}) {
	result := make(chan models.ResultChan[uuid.UUID], 1)
	committed := make(chan bool)
	insert := services.Service.Conne.DbInsert

	newDoc := models.Document{
		DocumentName: docPost.DocName,
		MongoID:      "test",
	}

	go func() {
		newDoc.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result

	go CreateDocContribution(docPost.UserUuid, insertResult.Data, committed)
	go CreateDocOwner(docPost.UserUuid, insertResult.Data, committed)

	// Loop to wait for signals from committed channel
	for i := 0; i < 2; i++ { // Expecting two signals for contributions and owner creation
		fmt.Println(<-committed)
	}

	if insertResult.Error != nil {
		log.Printf("Error inserting document: %v", insertResult.Error)
		res <- uuid.Nil
		close(res)
		return
	}

	res <- insertResult.Data
	defer close(res)
}
