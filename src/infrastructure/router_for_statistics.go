package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForStatistics(g *echo.Group, sqlHandler SqlHandler) {
	g.GET("/city", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).GetCityDataByCityId(c)
	})

	g.GET("/pref", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).GetCityDataByPrefId(c)
	})

	g.GET("/city/ranking/build_count", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).GetCityDataRanking(c)
	})

	g.GET("/city/ranking/build_count/new", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).(c)
	})

	g.GET("/pref/ranking/build_count", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataRanking(c)
	})
}
