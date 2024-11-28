package docs

import (
	"docs/internal/models"
	"docs/internal/services"
	"log"
)

// CreateDocPublicEdit handles the insertion of a new public edit document.
func CreateDocPublicEdit(documentEdit models.DocumentPublicEdit, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newDocEdit := models.DocumentPublicEdit(documentEdit)

	result := make(chan models.ResultChan[string], 1)

	go func() {
		newDocEdit.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting public edit document: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}

// UpdateDocPublicEdit handles the update of an existing public edit document.
func UpdateDocPublicEdit(documentEdit models.DocumentPublicEdit, res chan<- bool) {
	update := services.Service.Conne.DbUpdate
	newDocEdit := models.DocumentPublicEdit(documentEdit)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocEdit.Update(update, result)
		close(result)
	}()

	updateResult := <-result
	if updateResult.Error != nil {
		log.Printf("Error updating public edit document: %v", updateResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// DeleteDocPublicEdit handles the deletion of a public edit document.
func DeleteDocPublicEdit(documentEdit models.DocumentPublicEdit, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	newDocEdit := models.DocumentPublicEdit(documentEdit)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocEdit.Delete(delete, result)
		close(result)
	}()

	deleteResult := <-result
	if deleteResult.Error != nil {
		log.Printf("Error deleting public edit document: %v", deleteResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// QueryDocPublicEdit handles querying a public edit document.
func QueryDocPublicEdit(documentEdit models.DocumentPublicEdit, res chan<- *models.DocumentPublicEdit) {
	query := services.Service.Conne.DbRead
	newDocEdit := models.DocumentPublicEdit(documentEdit)

	result := make(chan models.ResultChan[*models.DocumentPublicEdit])

	go func() {
		newDocEdit.Query(query, result)
		close(result)
	}()

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying public edit document: %v", queryResult.Error)
		res <- nil
		return
	}
	res <- queryResult.Data
	defer close(res)
	defer close(result)
}
func CreateDocPublicView(documentView models.DocumentPublicView, res chan<- bool) {
	insert := services.Service.Conne.DbInsert
	newDocView := models.DocumentPublicView(documentView)

	result := make(chan models.ResultChan[string], 1)

	go func() {
		newDocView.Insert(insert, result)
		close(result)
	}()

	insertResult := <-result
	if insertResult.Error != nil {
		log.Printf("Error inserting public view document: %v", insertResult.Error)
		res <- false
		return
	}
	res <- true
}

// UpdateDocPublicView handles the update of an existing public view document.
func UpdateDocPublicView(documentView models.DocumentPublicView, res chan<- bool) {
	update := services.Service.Conne.DbUpdate
	newDocView := models.DocumentPublicView(documentView)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocView.Update(update, result)
		close(result)
	}()

	updateResult := <-result
	if updateResult.Error != nil {
		log.Printf("Error updating public view document: %v", updateResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// DeleteDocPublicView handles the deletion of a public view document.
func DeleteDocPublicView(documentView models.DocumentPublicView, res chan<- bool) {
	delete := services.Service.Conne.DbDelete
	newDocView := models.DocumentPublicView(documentView)

	result := make(chan models.ResultChan[error])

	go func() {
		newDocView.Delete(delete, result)
		close(result)
	}()

	deleteResult := <-result
	if deleteResult.Error != nil {
		log.Printf("Error deleting public view document: %v", deleteResult.Error)
		res <- false
		return
	}
	res <- true
	defer close(res)
	defer close(result)
}

// QueryDocPublicView handles querying a public view document.
func QueryDocPublicView(documentView models.DocumentPublicView, res chan<- *models.DocumentPublicView) {
	query := services.Service.Conne.DbRead
	newDocView := models.DocumentPublicView(documentView)

	result := make(chan models.ResultChan[*models.DocumentPublicView])

	go func() {
		newDocView.Query(query, result)
		close(result)
	}()

	queryResult := <-result
	if queryResult.Error != nil {
		log.Printf("Error querying public view document: %v", queryResult.Error)
		res <- nil
		return
	}
	res <- queryResult.Data
	defer close(res)
	defer close(result)
}
