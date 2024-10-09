package sessions

import (
	"docs/internal/models"
	"docs/internal/services"
)

func GetSession(token string, res chan<- models.Session) {
	query := services.Service.Conne.DbRead
	session := models.Session{
		Token: token,
	}

	resultChan := make(chan models.ResultChan[models.Session], 1)

	go session.Query(query, resultChan)

	result := <-resultChan
	if result.Error != nil {
		res <- models.Session{}
		return
	}
	res <- result.Data
	defer close(res)
	defer close(resultChan)

}
