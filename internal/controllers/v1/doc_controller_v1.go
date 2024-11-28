package v1

import (
	dto "docs/internal/Dto"
	"docs/internal/response"
	docs "docs/internal/services/doc"
	"docs/internal/utils"
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
	userIdCookie, _ := c.Cookie("doc")
	var newDoc dto.NewDoc
	if err := c.ShouldBindJSON(&newDoc); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			BaseResponse: response.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request data",
			},
			Error: err.Error(),
		})
		return
	}
	userId, _ := utils.Deobfuscate(userIdCookie)
	docPost := dto.DocPost{
		UserUuid: uuid.MustParse(userId),
		DocName:  newDoc.Title,
	}
	result := make(chan interface{})
	go docs.CreateDoc(docPost, newDoc.Public, result)

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

// @Summary Deletes a document
// @Description Deletes a document by ID
// @Tags Document
// @Accept  json
// @Produce  json
// @Param   doc_id path string true "Document ID"
// @Success 200 {object} response.SuccessResponse "Document deleted successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request data"
// @Failure 500 {object} response.ErrorResponse "Internal server error "
// @Router /doc/{doc_id} [delete]
func DeleteDoc(c *gin.Context) {
	userIdCookie, _ := c.Cookie("doc")
	docId := c.Param("doc_id")
	userId, _ := utils.Deobfuscate(userIdCookie)
	docIdUuid, _ := uuid.Parse(docId)
	userDoc := docs.UserDoc{
		UserID:     uuid.MustParse(userId),
		DocumentID: docIdUuid,
	}
	result := make(chan bool)
	go docs.DeleteDoc(userDoc, result)

	res := <-result
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
		return
	}

	successfully := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Document deleted successfully",
		},
	}
	c.JSON(http.StatusOK, successfully)
}
