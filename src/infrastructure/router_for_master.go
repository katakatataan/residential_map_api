package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForMaster(g *echo.Group, sqlHandler SqlHandler) {
	g.GET("/prefcities", func(c echo.Context) error {
		return controller.NewMstPrefCityController(&sqlHandler).GetMstPrefCity(c)
	})
}
