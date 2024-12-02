package docs

import (
	"context"
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

	newDocContribution.Insert(insert, result)
	close(result)

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting document ownership: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}

func UpdateDocOwner(documentOwner models.DocumentOwnership, res chan<- bool) {
	update := services.Service.Conne.DbUpdate
	newDocOwner := models.DocumentOwnership(documentOwner)

	result := make(chan models.ResultChan[error])

	newDocOwner.Update(update, result)
	close(result)

	updateResult := <-result
	if updateResult.Error != nil {
		log.Printf("Error updating document owner: %v", updateResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

func DeleteDocOwner(documentOwner models.DocumentOwnership, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	newDocOwner := models.DocumentOwnership(documentOwner)

	result := make(chan models.ResultChan[error])

	newDocOwner.Delete(delete, result)
	close(result)

	deleteResult := <-result
	if deleteResult.Error != nil {
		log.Printf("Error deleting document owner: %v", deleteResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

func QueryDocOwner(documentOwner models.DocumentOwnership, res chan<- *models.DocumentOwnership) {
	query := services.Service.Conne.DbRead
	newDocOwner := models.DocumentOwnership(documentOwner)

	result := make(chan models.ResultChan[*models.DocumentOwnership])

	newDocOwner.Query(query, result)
	close(result)

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying document owner: %v", queryResult.Error)
		res <- nil
		return
	}
	res <- queryResult.Data
	defer close(res)
	defer close(result)
}

func QueryAllDocOwnerByUser(userID uuid.UUID, res chan<- *[]uuid.UUID) {
	query := services.Service.Conne.DbRead
	resultChan := make(chan models.ResultChan[[]uuid.UUID])

	// Create user ownership object
	user := models.DocumentOwnership{UserID: userID}

	// Query for document IDs asynchronously
	go user.QueryAllByUser(context.Background(), query, resultChan)

	// Receive the result
	queryResult := <-resultChan
	if queryResult.Error != nil {
		// Handle error - send nil to response channel
		res <- nil
		return
	}

	docIDs := queryResult.Data
	if docIDs == nil {
		// Handle case where no documents were found
		res <- &[]uuid.UUID{}
		return
	}

	// Simulate transforming UUIDs to Document objects
	var docs []uuid.UUID
	for _, docID := range docIDs {
		docs = append(docs, docID)
	}

	// Send result back
	res <- &docs
}
