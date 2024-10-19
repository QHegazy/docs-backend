package v1

import (
	dto "docs/internal/Dto"
	"docs/internal/response"
	docs "docs/internal/services/doc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Creates a new document
// @Description Takes in document data and creates a new document
// @Tags Document
// @Accept  json
// @Produce  json
// @Param   docPost body dto.DocPost true "Document Post Data"
// @Success 200 {object} response.SuccessResponse{data=map[string]interface{}}
// @Failure 400 {object} response.ErrorResponse "Invalid request data"
// @Failure 500 {object} response.ErrorResponse "Internal server error "
// @Router /doc [post]
func NewDoc(c *gin.Context) {
	var docPost dto.DocPost
	if err := c.ShouldBindJSON(&docPost); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			BaseResponse: response.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request data",
			},
			Error: err.Error(),
		})
		return
	}

	result := make(chan interface{})
	go docs.CreateDoc(docPost, result)

	res := <-result
	if res == uuid.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Document created successfully",
		},
		Data: map[string]interface{}{
			"mongoID": res,
		},
	}
	c.JSON(http.StatusOK, successfully)
}

// @Summary Retrieves all documents
// @Description Fetches all documents in the database
// @Tags Document
// @Accept  json
// @Produce  json
// @Success 200 {object} response.SuccessResponse "Documents retrieved successfully"
// @Router /doc [get]
func RetrieveDocs(c *gin.Context) {
	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Documents retrieved successfully",
		},
	}
	c.JSON(http.StatusOK, successfully)
}
