package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"

	"github.com/google/uuid"
)

// CreateBlock handles the insertion of a new block record.
func CreateBlock(block models.Block, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newBlock := models.Block(block)

	// Adjusting result channel type to match Insert return type
	result := make(chan models.ResultChan[uuid.UUID], 1)

	go func() {
		newBlock.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting block: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}

// UpdateBlock handles the update of an existing block record.
func UpdateBlock(block models.Block, res chan<- bool) {
	update := services.Service.Conne.DbUpdate
	newBlock := models.Block(block)

	result := make(chan models.ResultChan[error])

	go func() {
		newBlock.Update(update, result)
		close(result)
	}()

	updateResult := <-result
	if updateResult.Error != nil {
		log.Printf("Error updating block: %v", updateResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// DeleteBlock handles the deletion of a block record.
func DeleteBlock(block models.Block, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	newBlock := models.Block(block)

	result := make(chan models.ResultChan[error])

	go func() {
		newBlock.Delete(delete, result)
		close(result)
	}()

	deleteResult := <-result
	if deleteResult.Error != nil {
		log.Printf("Error deleting block: %v", deleteResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// QueryBlock handles querying a block record.
func QueryBlock(block models.Block, res chan<- *models.Block) {
	query := services.Service.Conne.DbRead
	newBlock := models.Block(block)

	result := make(chan models.ResultChan[*models.Block])

	go func() {
		newBlock.Query(query, result)
		close(result)
	}()

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying block: %v", queryResult.Error)
		res <- nil
		return
	}
	res <- queryResult.Data
	defer close(res)
	defer close(result)
}
