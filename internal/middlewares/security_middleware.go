package middlewares

import (
	"docs/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecurityMiddleware(expectedHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				BaseResponse: response.BaseResponse{
					Status:  http.StatusBadRequest,
					Message: "Bad Request",
				},
				Error: "Invalid Host",
			})
			return
		}

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")

		c.Next()
	}
}
