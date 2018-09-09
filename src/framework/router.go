package framework

import (
	"residential_map_api/src/interface/controller"
	"residential_map_api/src/usecase/interactor"
	"residential_map_api/src/usecase/repository"

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
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))

	sqlHandler := NewSqlHandler()
	prefCityController := controller.NewMstPrefCityController(
		interactor.NewMstPrefCityInteractor(
			repository.NewMstPrefCityRepository(&sqlHandler)))

	e.GET("/prefcities", func(c echo.Context) error {
		return prefCityController.GetMstPrefCity(c)
	})

	echopprof.Wrap(e)
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
