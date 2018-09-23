package infrastructure

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{Validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// validate項目に"BuildDate"タグを追加
	cv.Validator.RegisterValidation("BuildDate", DateValidator)
	return cv.Validator.Struct(i)
}

func DateValidator(fl validator.FieldLevel) bool {
	return fl.Field().String() == "test"
}
