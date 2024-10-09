package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"

	"github.com/google/uuid"
)

func CreateDoc(res chan<- interface{}) {
	result := make(chan models.ResultChan[uuid.UUID], 1)
	insert := services.Service.Conne.DbInsert

	newDoc := models.Document{
		DocumentName: uuid.NewString(),
		MongoID:      uuid.NewString(),
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

	res <- insertResult.Data
	defer close(res)

}
