package infrastructure

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

const timeFormat = "2006-01-02"

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("can-be-time", ValidateTimeString)
	return cv.validator.Struct(i)
}
func ValidateTimeString(fl validator.FieldLevel) bool {
	t, err := time.Parse(timeFormat, fl.Field().String())
	if err != nil && t.Unix() < 0 {
		return false
	}
	return true
}
