package infrastructure

import (
	"reflect"
	"strconv"

	"github.com/labstack/echo"
)

type CustomBinder struct{}

func NewBinder() (cb *CustomBinder) {
	return &CustomBinder{}
}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	// if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
	err = db.Bind(i, c)
	// pp.Println(reflect.ValueOf(i).Elem().FieldByName("From"))
	// pp.Println(reflect.ValueOf(i).Elem().FieldByName("To"))
	// 要件はparamタグに対して、適切なバリデーションをかける
	// if err != nil {
	// 	err = setTimeField(c.QueryParam("from"), reflect.ValueOf(i).Elem().FieldByName("From"))
	// 	err = setTimeField(c.QueryParam("to"), reflect.ValueOf(i).Elem().FieldByName("To"))
	// }
	//ここでtagでparamがきた場合にbindする
	// pathNames が回せるならjj
	// pathパラメータを取得してbindする。数字のみ必要であれば修正する
	for _, name := range c.ParamNames() {
		value, _ := strconv.Atoi(c.Param(name))
		bindPathParamsToStruct(value, reflect.ValueOf(i).Elem().FieldByName(name))
	}

	return err
}

func bindPathParamsToStruct(value int, field reflect.Value) {
	field.Set(reflect.ValueOf(value))
}

// func setTimeField(value string, field reflect.Value) error {
// 	if value == "" {
// 		value = "2018-01-01"
// 	}
// 	// boolVal, err := strconv.ParseBool(value)
// 	t, err := time.Parse(timeFormat, value)

// 	// pp.Println(t)
// 	// pp.Println(value)
// 	if err == nil {
// 		field.Set(reflect.ValueOf(t))
// 	}
// 	return err
// }
