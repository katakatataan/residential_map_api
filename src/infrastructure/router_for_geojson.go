package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForGeojson(g *echo.Group, sqlHandler SqlHandler) {
	g.GET("/pref", func(c echo.Context) error {
		return controller.NewGeoPrefectureController(&sqlHandler).GeoPlainPrefecture(c)
	})
}
