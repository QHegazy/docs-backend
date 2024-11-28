package dto

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Visibility int

const (
	View Visibility = iota + 1
	Edit
	Private
)

var visibilityStrings = map[Visibility]string{
	View: "view",
	Edit: "edit",
	Private:  "private",
}

func (v Visibility) String() string {
	return visibilityStrings[v]
}

func ParseVisibility(s string) (Visibility, error) {
	for k, v := range visibilityStrings {
		if v == s {
			return k, nil
		}
	}
	return 0, errors.New("invalid visibility")
}

func (v *Visibility) UnmarshalJSON(data []byte) error {
	var visibilityString string
	if err := json.Unmarshal(data, &visibilityString); err != nil {
		return err
	}
	visibility, err := ParseVisibility(visibilityString)
	if err != nil {
		return err
	}
	*v = visibility
	return nil
}

type DocPost struct {
	UserUuid uuid.UUID `form:"user_uuid" json:"user_uuid" binding:"required,uuid"`
	DocName  string    `form:"title" json:"title" binding:"required,min=3,max=50"`
}
type NewDoc struct {
	Title  string     `form:"title" json:"title" binding:"required,min=3,max=50"`
	Public Visibility `form:"public" json:"public" binding:"required"`
}

var validate = validator.New()

func init() {
	validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err == nil
	})

	validate.RegisterValidation("visibility", func(fl validator.FieldLevel) bool {
		val, ok := fl.Field().Interface().(Visibility)
		return ok && (val == View || val == Edit || val == Private)
	})
}
