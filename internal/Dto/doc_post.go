package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DocPost struct {
	UserUuid uuid.UUID `form:"user_uuid" json:"user_uuid" binding:"required,uuid"`
	DocName  string    `form:"name" json:"name" binding:"required,min=1,max=255"`
}

// Initialize a validator instance
var validate = validator.New()

// Register UUID validation
func init() {
	validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err == nil
	})
}
