package handlers

import (
	// your generated docs
	_ "docs/docs"
	v1 "docs/internal/controllers/v1"
	"docs/internal/middlewares"
	"os"

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
		v1Group.GET("/authorize", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.Authorize_v1)
		v1Group.POST("/logout", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.Logout_v1)
		v1Group.GET("/doc/:doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocs)
		v1Group.POST("/doc", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.NewDoc)
		// v1Group.DELETE("/doc",middlewares.AuthMiddleware(), middlewares.CheckSessionToken(),v1.)

		// v1Group.GET("/doc/:doc_id/contributors", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocContributors)
		// v1Group.POST("/doc/:doc_id/contributors", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.AddDocContributor)
		// v1Group.DELETE("/doc/:doc_id/contributors/:user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RemoveDocContributor)
		// v1Group.GET("/doc/:doc_id/owners", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocOwners)
		// v1Group.POST("/doc/:doc_id/owners", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.AddDocOwner)
		// v1Group.DELETE("/doc/:doc_id/owners/:user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RemoveDocOwner)
		// v1Group.GET("/doc/:doc_id/access", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocAccess)
		// v1Group.POST("/doc/:doc_id/access", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.SetDocAccess)
		// v1Group.DELETE("/doc/:doc_id/access", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RemoveDocAccess)
		// v1Group.GET("/user", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUser)
		// v1Group.GET("/user/:user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserById)
		// v1Group.GET("/user/:user_id/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserDocs)
		// v1Group.GET("/user/:user_id/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserBlocked)
		// v1Group.POST("/user/:user_id/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.BlockUser)
		// v1Group.DELETE("/user/:user_id/blocked/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.UnblockUser)
		// v1Group.GET("/user/:user_id/blocked/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckUserBlocked)
		// v1Group.GET("/user/:user_id/blocked/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserBlockedDocs)
		// v1Group.GET("/user/:user_id/blocked/docs/:doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserBlockedDoc)
		// v1Group.POST("/user/:user_id/blocked/docs/:doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.BlockDoc)
		// v1Group.DELETE("/user/:user_id/blocked/docs/:doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.UnblockDoc)
		// v1Group.GET("/user/:user_id/blocked/docs/:doc_id/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckDocBlocked)
		// v1Group.GET("/user/:user_id/blocked/docs/:doc_id/blocked/users", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveDocBlockedUsers)
		// v1Group.GET("/user/:user_id/blocked/docs/:doc_id/blocked/users/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckDocBlockedUser)
		// v1Group.POST("/user/:user_id/blocked/docs/:doc_id/blocked/users/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.BlockDocUser)
		// v1Group.DELETE("/user/:user_id/blocked/docs/:doc_id/blocked/users/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.UnblockDocUser)
		// v1Group.GET("/user/:user_id/sessions", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessions)
		// v1Group.POST("/user/:user_id/sessions", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CreateUserSession)
		// v1Group.DELETE("/user/:user_id/sessions/:session_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.DeleteUserSession)
		// v1Group.GET("/user/:user_id/sessions/:session_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSession)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUser)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserDocs)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserBlocked)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserBlockedDocs)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserBlockedDoc)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckUserSessionUserDocBlocked)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id)
		// /blocked/users", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserDocBlockedUsers)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckUserSessionUserDocBlockedUser)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id/docs", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserDocBlockedUserDocs)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id/docs/:blocked_doc_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserDocBlockedUserDoc)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id/docs/:blocked_doc_id/blocked", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckUserSessionUserDocBlockedUserDocBlocked)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id/docs/:blocked_doc_id/blocked/users", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.RetrieveUserSessionUserDocBlockedUserDocBlockedUsers)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:doc_id
		// /blocked/users/:blocked_user_id/docs/:blocked_doc_id/blocked/users/:blocked_user_id", middlewares.AuthMiddleware(), middlewares.CheckSessionToken(), v1.CheckUserSessionUserDocBlockedUserDocBlockedUser)
		// v1Group.GET("/user/:user_id/sessions/:session_id/user/blocked/docs/:")
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
