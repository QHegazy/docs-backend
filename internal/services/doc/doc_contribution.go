package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"
)

func CreateDocContribution(userId string, docId string, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newDocContribution := models.DocumentContribution{
		UserID:     userId,
		DocumentID: docId,
	}

	result := make(chan models.ResultChan[string])

	go func() {
		newDocContribution.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting document: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}
