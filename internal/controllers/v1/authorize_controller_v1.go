package v1

import (
	"docs/internal/models"
	"docs/internal/response"
	"docs/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Name     string `json:"Name"`
	ImageURL string `json:"ImageURL"`
	Email    string `json:"Email"`
}

func Authorize_v1(c *gin.Context) {
	cookieValue, _ := c.Cookie("lg")

	resultchan := make(chan models.ResultChan[models.User])

	go auth.Authorize(cookieValue, resultchan)

	result := <-resultchan

	// Create a new instance of UserResponse
	userResponse := UserResponse{
		Name:     result.Data.Name,
		ImageURL: result.Data.ImageURL,
		Email:    result.Data.Email,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Authorized",
		},
		Data: userResponse,
	})
}

func Logout_v1(c *gin.Context) {
	resultchan := make(chan models.ResultChan[models.Session])

	go auth.Logout(c, resultchan)

	c.JSON(http.StatusOK, response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Logged out",
		},
		Data: <-resultchan,
	})
}
