package handlers

import (
	v1 "docs/internal/controllers/v1"
	"docs/internal/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	expectedHost := os.Getenv("HOST")
	r.Use(middlewares.InternalServerErrorMiddleware(), middlewares.SecurityMiddleware(expectedHost))
	r.NoRoute(middlewares.NotFound)
	v1Group := r.Group("v1")
	{
		v1Group.GET("/auth/google/login", v1.GoogleAuth)
		v1Group.GET("/auth/google/callback", v1.GoogleAuthCallback)
		v1Group.GET("/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocs)
		v1Group.POST("/new-doc", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.NewDoc)
	}

	return r
}
