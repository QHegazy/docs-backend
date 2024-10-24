package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"

	"github.com/google/uuid"
)

func CreateDocOwner(userId uuid.UUID, docId uuid.UUID, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newDocContribution := models.DocumentOwnership{
		UserID:     userId,
		DocumentID: docId,
	}

	result := make(chan models.ResultChan[string], 1)

	go func() {
		newDocContribution.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting document ownership: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}
