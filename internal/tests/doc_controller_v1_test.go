package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "docs/internal/controllers/v1"
	"docs/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewDoc_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/docs", v1.NewDoc)

	// Create the HTTP POST request with invalid data
	req, _ := http.NewRequest(http.MethodPost, "/docs", bytes.NewBuffer([]byte(`{"user_uuid": "invalid-uuid", "name": ""}`)))
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var responseBody response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid request data", responseBody.Message)
	assert.Contains(t, responseBody.Error, "invalid")
}

func TestRetrieveDocs_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/docs", v1.RetrieveDocs)

	// Create the HTTP GET request
	req, _ := http.NewRequest(http.MethodGet, "/docs", nil)

	// Send the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody response.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Documents retrieved successfully", responseBody.Message)
}
