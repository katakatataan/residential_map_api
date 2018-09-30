package infrastructure

import (
	"reflect"
	"time"

	"github.com/labstack/echo"
)

const timeFormat = "2006-01-02"

type CustomBinder struct{}

func NewBinder() (cb *CustomBinder) {
	return &CustomBinder{}
}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	// if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
	err = db.Bind(i, c)
	if err == nil {
		return
	}
	// pp.Println(reflect.ValueOf(i).Elem().FieldByName("From"))
	// pp.Println(reflect.ValueOf(i).Elem().FieldByName("To"))
	if err != nil {
		err = setTimeField(c.QueryParam("from"), reflect.ValueOf(i).Elem().FieldByName("From"))
		err = setTimeField(c.QueryParam("to"), reflect.ValueOf(i).Elem().FieldByName("To"))
	}

	return err
}
func setTimeField(value string, field reflect.Value) error {
	if value == "" {
		value = "2018-01-01"
	}
	// boolVal, err := strconv.ParseBool(value)
	t, err := time.Parse(timeFormat, value)

	// pp.Println(t)
	// pp.Println(value)
	if err == nil {
		field.Set(reflect.ValueOf(t))
	}
	return err
}
