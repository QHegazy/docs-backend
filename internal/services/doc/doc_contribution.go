package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"
)

func CreateDocContribution(document_contribution models.DocumentContribution, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newDocContribution := models.DocumentContribution(document_contribution)

	result := make(chan models.ResultChan[string], 1)

	go func() {
		newDocContribution.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting document contribution: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}
func UpdateDocContribution(document_contribution models.DocumentContribution, res chan<- bool) {
	update := services.Service.Conne.DbUpdate
	newDocContribution := models.DocumentContribution(document_contribution)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocContribution.Update(update, result)
		close(result)
	}()

	updateResult := <-result
	if updateResult.Error != nil {
		log.Printf("Error updating document: %v", updateResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

func DeleteDocContribution(document_contribution models.DocumentContribution, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	newDocContribution := models.DocumentContribution(document_contribution)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocContribution.Delete(delete, result)
		close(result)
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

func QueryDocContribution(document_contribution models.DocumentContribution, res chan<- *models.DocumentContribution) {
	query := services.Service.Conne.DbRead
	newDocContribution := models.DocumentContribution(document_contribution)

	result := make(chan models.ResultChan[*models.DocumentContribution])

	go func() {
		newDocContribution.Query(query, result)
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
