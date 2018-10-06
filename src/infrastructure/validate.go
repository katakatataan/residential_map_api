package infrastructure

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

const timeFormat = "2006-01-02"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{Validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterValidation("canbetime", CanBeTime)
	return cv.Validator.Struct(i)
}

func CanBeTime(fl validator.FieldLevel) bool {
	pp.Println(fl.Field().String())
	t, err := time.Parse(timeFormat, fl.Field().String())
	pp.Println(t)
	if err != nil {
		return true
	}
	return true
}
