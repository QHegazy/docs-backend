package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DocPost struct {
	UserUuid uuid.UUID `form:"user_uuid" json:"user_uuid" binding:"required,uuid"`
	DocName  string    `form:"title" json:"title" binding:"required,min=3,max=50"`
}
type NewDoc struct {
	Title string `form:"title" json:"title" binding:"required,min=3,max=50"`
}

var validate = validator.New()

func init() {
	validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err == nil
	})
}
