package infrastructure

import (
	"residential_map_api/src/interface/controller"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/sevenNt/echo-pprof"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Validator
	e.Validator = NewValidator()

	// Bind
	// e.Binder = NewBinder()
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))

	sqlHandler := NewSqlHandler()
	mstPrefCityController := controller.NewMstPrefCityController(&sqlHandler)
	cityDataController := controller.NewCityDataController(&sqlHandler)
	geoPrefectureController := controller.NewGeoPrefectureController(&sqlHandler)
	e.GET("/mst_prefcities", func(c echo.Context) error {
		return mstPrefCityController.GetMstPrefCity(c)
	})

	e.GET("/citydata", func(c echo.Context) error {
		return cityDataController.GetCityDataByCityId(c)
	})

	e.GET("/prefdata", func(c echo.Context) error {
		return cityDataController.GetCityDataByPrefId(c)
	})

	e.GET("/monthlycitydatarank", func(c echo.Context) error {
		return cityDataController.GetCityDataRanking(c)
	})

	e.GET("/monthlyprefdatarank", func(c echo.Context) error {
		return cityDataController.GetPrefDataRanking(c)
	})
	e.GET("/geopref", func(c echo.Context) error {
		return geoPrefectureController.GeoPlainPrefecture(c)
	})

	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":1323"))
}
