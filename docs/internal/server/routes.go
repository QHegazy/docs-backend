package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.GET("/", s.HelloWorldHandler)
    r.GET("auth/google/login",s.googleAuth)
    r.GET("auth/google/callback",s.googleAuthCallback)
    clientId:=os.Getenv("GOOGLE_CLIENT_ID")
    clientSecret:=os.Getenv("GOOGLE_CLIENT_SECRET")
    callbackUrl:=os.Getenv("GOOGLE_CALLBACK_URL")

	r.GET("/health", s.healthHandler)
	goth.UseProviders(
		google.New(
            clientId,
			clientSecret,
			callbackUrl,
			"email", 
			"profile",
		),
	)
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) googleAuth(c *gin.Context){
    c.Request.URL.RawQuery = "provider=google"
    gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (s *Server) googleAuthCallback(c *gin.Context){
    user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "user":     user,
            "provider": "google",
        },
    })
    

}