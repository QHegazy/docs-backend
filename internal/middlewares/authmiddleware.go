package middlewares

import (
	"docs/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("lg")
		if err != nil || cookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				BaseResponse: response.BaseResponse{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				},
				Error: "Session token is missing or invalid.",
			})
			return
		}

		c.Next()
	}
}
