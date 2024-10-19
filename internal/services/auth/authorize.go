package auth

import (
	"docs/internal/database"
	"docs/internal/models"
	"docs/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
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

func Logout(c *gin.Context, resultchan chan<- models.ResultChan[models.Session]) {
	cookieValue, _ := c.Cookie("lg")
	delete := services.Service.Conne.DbDelete
	session := models.Session{
		Token: cookieValue,
	}
	result := make(chan models.ResultChan[models.Session], 1)
	go func() {
		session.DeleteByToken(delete, result)
		c.SetCookie("lg", "", -1, "/", "", false, true)
		resultchan <- <-result
		
		defer close(result)
		defer close(resultchan)
	}()

}
