package auth

import (
	"docs/internal/models"
	"docs/internal/services"
	"fmt"

	"github.com/markbates/goth"
)

func Register(user *goth.User) {

	insert := services.Service.Conne.DbInsert
	newUser := models.User{
		Name:     user.Name,
		OauthID:  user.UserID,
		ImageURL: user.AvatarURL,
		Email:    user.Email,
	}

	go func() {
		if err := newUser.Insert(insert); err != nil {
			// Log the error if insert fails
			fmt.Printf("Failed to insert user: %v\n", err)
		}
	}()

}

// func Login(user models.User) {
// 	insert := services.Service.Conne.DbInsert
// 	newUser := models.User{
// 		Name:     user.Name,
// 		OauthID:  user.OauthID,
// 		ImageURL: user.ImageURL,
// 		Email:    user.Email,
// 	}
// 	if err := newUser.Insert(insert); err != nil {
// 		// Log the error if insert fails
// 		fmt.Printf("Failed to insert user: %v\n", err)
// 	}

// }
