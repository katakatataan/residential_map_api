package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForPref(g *echo.Group, sqlHandler SqlHandler) {
	// 表から未使用
	g.GET("/:PrefId", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataByPrefId(c)
	})

	// 表から未使用
	g.GET("/ranking/build_count", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataRankingBuildCount(c)
	})
}
