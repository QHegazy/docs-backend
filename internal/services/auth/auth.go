package auth

import (
	"docs/internal/models"
	"docs/internal/services"
	"docs/internal/utils"
	"fmt"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

// Login handles user login by checking OAuth ID and creating a session.
func Login(user *goth.User, token chan<- string) {
	userToken := make(chan string)
	userID := checkUserByOauthID(user.UserID)

	if userID != uuid.Nil {
		// Existing user: create a session
		go CreateSession(userID, userToken)
	} else {
		// New user: insert into the database
		insert := services.Service.Conne.DbInsert
		newUser := models.User{
			Name:     user.Name,
			OauthID:  user.UserID,
			ImageURL: user.AvatarURL,
			Email:    user.Email,
		}

		userIDChan := make(chan models.ResultChan[uuid.UUID])
		go newUser.Insert(insert, userIDChan)

		// Wait for the result from the insert operation
		result := <-userIDChan
		if result.Error != nil {
			token <- "" // Send empty token on error
			fmt.Printf("Error inserting new user: %v\n", result.Error)
			return
		}

		go CreateSession(result.Data, userToken) // Create a session for the new user
	}

	go func() {
		sessionToken := <-userToken
		token <- sessionToken
	}()
}

// CreateSession generates a session token and inserts a new session into the database.
func CreateSession(userID uuid.UUID, userToken chan<- string) {
	token, err := utils.GenerateToken(128)
	if err != nil {
		userToken <- "" // Send empty token on error
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

	// Wait for the result from Insert
	result := <-resultChan
	if result.Error != nil {
		userToken <- "" // Send empty token on error
		fmt.Printf("Failed to insert session: %v\n", result.Error)
		return
	}

	userToken <- token // Send the generated token
}

// checkUserByOauthID checks if a user with the given OAuth ID exists in the database.
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
