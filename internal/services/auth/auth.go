package auth

import (
	"docs/internal/models"
	"docs/internal/services"
	"docs/internal/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

func Login(user *goth.User, token chan<- string) {
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
		user_Id, err := newUser.Insert(insert)
		if err != nil {
			token <- ""
			fmt.Printf("Error inserting new user: %v\n", err)
			return
		}

		go CreateSession(user_Id, userToken)
	}

	go func() {
		sessionToken := <-userToken
		token <- sessionToken
	}()

}

func CreateSession(UserID uuid.UUID, userToken chan<- string) {
	token, err := utils.GenerateToken(32)
	if err != nil {
		userToken <- ""
		fmt.Printf("Failed to generate token: %v\n", err)
		return
	}

	newSession := models.Session{
		UserID:    UserID,
		Token:     token,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Online:    true,
	}

	insert := services.Service.Conne.DbInsert
	if err := newSession.Insert(insert); err != nil {
		userToken <- ""
		fmt.Printf("Failed to insert session: %v\n", err)
		return
	}

	userToken <- token
}

func checkUserByOauthID(oauthID string) uuid.UUID {
	ch := make(chan uuid.UUID)
	query := services.Service.Conne.DbRead
	newUser := models.User{}
	go newUser.UserIdQuery(query, oauthID, ch)
	userID := <-ch

	if userID == uuid.Nil {
		return uuid.Nil
	}
	return userID
}
