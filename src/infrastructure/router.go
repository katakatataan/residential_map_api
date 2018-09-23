package infrastructure

import (
	"residential_map_api/src/interface/controller"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/sevenNt/echo-pprof"

	_ "github.com/lib/pq"
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
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))

	sqlHandler := NewSqlHandler()
	mstPrefCityController := controller.NewMstPrefCityController(&sqlHandler)
	cityDataController := controller.NewCityDataController(&sqlHandler)

	e.GET("/prefcities", func(c echo.Context) error {
		return mstPrefCityController.GetMstPrefCity(c)
	})

	e.GET("/citydata", func(c echo.Context) error {
		return cityDataController.GetCityData(c)
	})
	e.GET("/citydata/:id", func(c echo.Context) error {
		return cityDataController.GetCityDataById(c)
	})

	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":1323"))
}
