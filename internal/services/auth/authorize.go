package auth

import (
	"docs/internal/database"
	"docs/internal/models"
	"fmt"
)

func Authorize(cookieValue string, resultchan chan<- models.ResultChan[models.User]) {
	query := database.New().DbRead
	result := make(chan models.ResultChan[models.User], 1) // Result channel for the query

	session := models.Session{
		Token: cookieValue,
	}
	fmt.Println(session)
	go func() {
		session.QueryUserId(query, result)
		resultchan <- <-result
		defer close(result)
		defer close(resultchan)
	}()
}
