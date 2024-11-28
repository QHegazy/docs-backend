package auth

import (
	"docs/internal/models"
	"docs/internal/services"
	"docs/internal/utils"
	"fmt"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

type UserAuth struct {
	UserID string `json:"Name"`
	Token  string `json:"Token"`
}

func Login(user *goth.User, userAuth chan<- UserAuth) {
	userToken := make(chan string)
	userID := checkUserByOauthID(user.UserID)
	if userID != uuid.Nil {
		go CreateSession(userID, userToken)
	} else {
		insert := services.Service.Conne.DbInsert
		newUser := models.User{
			Name:     user.Name,
			OauthID:  user.UserID,
			ImageURL: user.AvatarURL,
			Email:    user.Email,
		}

		userIDChan := make(chan models.ResultChan[uuid.UUID])
		go newUser.Insert(insert, userIDChan)

		result := <-userIDChan
		if result.Error != nil {
			userAuth <- UserAuth{}
			fmt.Printf("Error inserting new user: %v\n", result.Error)
			return
		}
		userID = result.Data
		go CreateSession(userID, userToken)
	}

	go func() {
		sessionToken := <-userToken
		if sessionToken == "" {
			userAuth <- UserAuth{}
			return
		}
		userAuth <- UserAuth{
			UserID: userID.String(),
			Token:  sessionToken,
		}
	}()
}

func CreateSession(userID uuid.UUID, userToken chan<- string) {
	token, err := utils.GenerateToken(128)
	if err != nil {
		userToken <- ""
		fmt.Printf("Failed to generate token: %v\n", err)
		return
	}

	newSession := models.Session{
		UserID:    userID,
		Token:     token,
		ExpiresAt: utils.GenerateExpireDate(7),
		Online:    true,
	}

	insert := services.Service.Conne.DbInsert
	resultChan := make(chan models.ResultChan[models.Session])
	go newSession.Insert(insert, resultChan)

	result := <-resultChan
	if result.Error != nil {
		userToken <- ""
		fmt.Printf("Failed to insert session: %v\n", result.Error)
		return
	}

	userToken <- token
}

func checkUserByOauthID(oauthID string) uuid.UUID {
	ch := make(chan uuid.UUID)
	query := services.Service.Conne.DbRead
	newUser := models.User{}

	go newUser.UserOauthIdQuery(query, oauthID, ch)
	userID := <-ch

	if userID == uuid.Nil {
		return uuid.Nil
	}
	return userID
}
