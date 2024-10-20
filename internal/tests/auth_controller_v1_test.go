package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "docs/internal/controllers/v1"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/auth/google", v1.GoogleAuth)

	req, _ := http.NewRequest(http.MethodGet, "/auth/google", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	assert.Contains(t, w.Header().Get("Location"), "https://accounts.google.com")
}

func TestGoogleAuthCallback(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/auth/google/callback", v1.GoogleAuthCallback)

	req, _ := http.NewRequest(http.MethodGet, "/auth/google/callback?code=some-auth-code", nil)
	w := httptest.NewRecorder()

	gothic.CompleteUserAuth = func(w http.ResponseWriter, r *http.Request) (goth.User, error) {
		return goth.User{
			UserID:   "1234s",
			Provider: "google",
			Name:     "John Doe",
			Email:    "john.doe@example.com",
		}, nil
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusPermanentRedirect, w.Code)
}
