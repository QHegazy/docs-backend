package server

import (
	"docs/internal/middlewares"
	"docs/internal/services/auth"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func oauth() {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("GOOGLE_CALLBACK_URL")
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
	r.GET("/", s.HelloWorldHandler)
	r.GET("auth/google/login", s.googleAuth)
	r.GET("auth/google/callback", s.googleAuthCallback)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	// hello := models.User{
	// 	Name:     uuid.NewString(),
	// 	OauthID:  uuid.NewString(),
	// 	ImageURL: uuid.NewString(),
	// 	Email:    uuid.NewString(),
	// }
	// auth.Login(hello)
	c.JSON(http.StatusOK, resp)
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
	go auth.Register(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":     user,
			"provider": "google",
		},
	})
}
