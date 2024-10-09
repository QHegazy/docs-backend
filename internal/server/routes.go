package server

import (
	"docs/internal/middlewares"
	"docs/internal/response"
	"docs/internal/services/auth"
	docs "docs/internal/services/doc"
	"docs/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func oauth() {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("GOOGLE_CALLBACK_URL")
	sessionSecret := os.Getenv(" SESSION_SECRET")
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
func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	expectedHost := os.Getenv("HOST")
	r.Use(middlewares.SecurityMiddleware(expectedHost))
	r.NoRoute(middlewares.NotFound)
	r.Use(middlewares.InternalServerErrorMiddleware())
	r.GET("auth/google/login", s.googleAuth)
	r.GET("auth/google/callback", s.googleAuthCallback)
	r.GET("/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), s.retraiveDocs)
	r.GET("/", s.homeApi)

	r.POST("newdoc", middlewares.AuthMiddleware(), s.newDoc)

	return r
}

func (s *Server) googleAuth(c *gin.Context) {
	oauth()
	c.Request.URL.RawQuery = "provider=google"
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (s *Server) googleAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := make(chan string)

	go func() {
		auth.Login(&user, token)
	}()
	select {
	case userToken := <-token:
		expireDate := utils.GenerateExpireDate(7)
		c.SetCookie("lg", userToken, int(expireDate.Unix()-time.Now().Unix()), "/", "", false, true)

		successfully := response.SuccessResponse{
			BaseResponse: response.BaseResponse{
				Status:  http.StatusOK,
				Message: "successfull login",
			},
		}
		c.SecureJSON(http.StatusOK, successfully)
	case <-time.After(2 * time.Second):
		c.SecureJSON(http.StatusGatewayTimeout, gin.H{"error": "Login request timed out"})
		return
	}

}

func (s *Server) retraiveDocs(c *gin.Context) {
	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "successfull login",
		},
	}
	c.JSON(http.StatusOK, successfully)

}

func (s *Server) homeApi(c *gin.Context) {
	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "successfull login",
		},
	}
	c.JSON(http.StatusOK, successfully)

}
func (s *Server) newDoc(c *gin.Context) {
	result := make(chan interface{})

	go docs.CreateDoc(result)

	res := <-result
	if res == uuid.Nil {
		return
	}
	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Document created successfully",
		},
		Data: map[string]interface{}{
			"mongoID": res,
		},
	}
	c.JSON(http.StatusOK, successfully)
}
