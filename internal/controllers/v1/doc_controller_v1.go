package v1

import (
	dto "docs/internal/Dto"
	"docs/internal/response"
	docs "docs/internal/services/doc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func RetrieveDocs(c *gin.Context) {
	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Documents retrieved successfully",
		},
	}
	c.JSON(http.StatusOK, successfully)
}
