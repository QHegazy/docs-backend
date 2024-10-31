package v1

import (
	"docs/internal/response"
	"docs/internal/services/auth"
	"docs/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func initOAuth() {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("GOOGLE_CALLBACK_URL")
	sessionSecret := os.Getenv("SESSION_SECRET")
	store := sessions.NewCookieStore([]byte(sessionSecret))

	gothic.Store = store
	goth.UseProviders(
		google.New(
			clientId,
			clientSecret,
			callbackUrl,
			"email",
			"profile",
		),
	)
}

func GoogleAuth(c *gin.Context) {
	initOAuth()
	c.Request.URL.RawQuery = "provider=google"
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GoogleAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := make(chan auth.UserAuth)

	go func() {
		auth.Login(&user, token)
	}()
	frontend := os.Getenv("FRONTENDREDIRECT")

	select {
	case userToken := <-token:
		userID := utils.Obfuscate(userToken.UserID)
		expireDate := utils.GenerateExpireDate(7)
		c.SetCookie("lg", userToken.Token, int(expireDate.Unix()-time.Now().Unix()), "/", "", true, true)
		c.SetCookie("doc", userID, int(expireDate.Unix()-time.Now().Unix()), "/v1", "", true, true)

		c.Redirect(http.StatusPermanentRedirect, frontend)
	case <-time.After(2 * time.Second):
		res := response.ErrorResponse{
			BaseResponse: response.BaseResponse{
				Status:  http.StatusGatewayTimeout,
				Message: "Login request timed out",
			},
			Error: "Login request timed out",
		}
		c.SecureJSON(http.StatusGatewayTimeout, res)
	}
}
