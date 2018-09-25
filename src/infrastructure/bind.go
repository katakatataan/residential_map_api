package infrastructure

import "github.com/labstack/echo"

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}
	//{TODO}ここでカスタムbindを定義する

	return
}
