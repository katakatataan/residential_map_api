package infrastructure

import (
	"residential_map_api/src/interface/controller"

	"github.com/labstack/echo"
)

func routeForPref(g *echo.Group, sqlHandler SqlHandler) {
	// もとはstatistics/pref
	g.GET("/:PrefId", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataByPrefId(c)
	})

	// もとはstatistics/pref/ranking/build_count
	g.GET("/ranking/build_count", func(c echo.Context) error {
		return controller.NewPrefDataController(&sqlHandler).GetPrefDataRankingBuildCount(c)
	})
}
