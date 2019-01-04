package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForStatistics(g *echo.Group, sqlHandler SqlHandler) {
	g.GET("/city", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).GetCityDataByCityId(c)
	})

	g.GET("/city/monthly", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).GetCityDataByTargetPeriod(c)
	})

	g.GET("/city/ranking/build_count/new", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).FindCityRankingBuildCount(c)
	})

	g.GET("/pref", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataByPrefId(c)
	})

	g.GET("/pref/ranking/build_count", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataRankingBuildCount(c)
	})
}
