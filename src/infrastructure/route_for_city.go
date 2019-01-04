package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForCity(g *echo.Group, sqlHandler SqlHandler) {
	// 表から未使用
	g.GET("/:CityId", func(c echo.Context) error {
		// 命名規則は一つの市区町村にたいしてなのでfindではなくget
		return controller.NewCityDataController(&sqlHandler).GetCityDataByCityId(c)
	})

	g.GET("/:CityId/monthly", func(c echo.Context) error {
		// 命名規則は一つの市区町村にたいしてなのでfindではなくget
		// 対象期間TargetPeriod
		return controller.NewCityDataController(&sqlHandler).GetCityDataByTargetPeriod(c)
	})

	// // もとはstatistics/city/ranking/build_count/new
	g.GET("/ranking/build_count", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).FindCityRankingBuildCount(c)
	})

	g.GET("/geojson/build_count", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).FindCitiesGeojsonWithBuildCount(c)
	})

	g.GET("/geojson", func(c echo.Context) error {
		return controller.NewCityDataController(&sqlHandler).FindCitiesGeojson(c)
	})
}
