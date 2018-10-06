package infrastructure

import (
	"time"

	"github.com/k0kubun/pp"
	"gopkg.in/go-playground/validator.v9"
)

const timeFormat = "2006-01-02"

// type CustomValidator struct {
// 	Validator *validator.Validate
// }

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func ValidateTimeString(fl validator.FieldLevel) bool {
	pp.Println(fl.Field().String())
	t, err := time.Parse(timeFormat, fl.Field().String())
	pp.Println(t)
	if err != nil {
		return true
	}
	return true
}
