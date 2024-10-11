// tests/auth_controller_v1_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "docs/internal/controllers/v1" // Adjust import path as necessary

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAuth(t *testing.T) {
	// Set up the Gin router for testing
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Define your route
	router.GET("/auth/google", v1.GoogleAuth)

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/auth/google", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the redirect status and location
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	assert.Contains(t, w.Header().Get("Location"), "https://accounts.google.com")
}

func TestGoogleAuthCallback(t *testing.T) {
	// Set up the Gin router for testing
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Define your route
	router.GET("/auth/google/callback", v1.GoogleAuthCallback)

	// Simulate a request with an authorization code
	req, _ := http.NewRequest(http.MethodGet, "/auth/google/callback?code=some-auth-code", nil)
	w := httptest.NewRecorder()

	// Mock the goth package's CompleteUserAuth method
	gothic.CompleteUserAuth = func(w http.ResponseWriter, r *http.Request) (goth.User, error) {
		// Return a mock user instead of calling the real OAuth process
		return goth.User{
			UserID:   "1234s",
			Provider: "google",
			Name:     "John Doe",
			Email:    "john.doe@example.com",
		}, nil
	}

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the expected response
	assert.Equal(t, http.StatusOK, w.Code)
}
