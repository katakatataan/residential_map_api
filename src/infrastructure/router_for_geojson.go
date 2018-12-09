package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForGeojson(g *echo.Group, sqlHandler SqlHandler) {
	// TODO: エンドポイントをprefなのかcityなのか再度検討する
	g.GET("/pref", func(c echo.Context) error {
		return controller.NewGeoPrefectureController(&sqlHandler).GeoPlainPrefecture(c)
	})
	g.GET("/city/build_count", func(c echo.Context) error {
		return controller.NewGeoPrefectureController(&sqlHandler).GeoCityBuildCount(c)
	})
}
