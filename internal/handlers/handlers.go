package handlers

import (
	// your generated docs
	v1 "docs/internal/controllers/v1"
	"docs/internal/middlewares"
	"os"
	_ "docs/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	expectedHost := os.Getenv("HOST")

	r.Use(middlewares.CORSMiddleware(), middlewares.InternalServerErrorMiddleware(), middlewares.SecurityMiddleware(expectedHost))
	r.NoRoute(middlewares.NotFound)
	r.GET("/auth/google/callback", v1.GoogleAuthCallback)
	v1Group := r.Group("v1")
	{

		v1Group.GET("/auth/google/login", v1.GoogleAuth)
		v1Group.GET("/doc", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocs)
		v1Group.POST("/doc", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.NewDoc)
		v1Group.GET("/authorize", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.Authorize_v1)
		v1Group.POST("/logout", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.Logout_v1)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
