package infrastructure

import (
	"os"

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
	e.Binder = NewBinder()

	sqlHandler := NewSqlHandler()
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))
	routeForCity(e.Group("/cities"), sqlHandler)
	routeForPref(e.Group("/prefs"), sqlHandler)
	routeForMaster(e.Group("/master"), sqlHandler)

	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":1323"))
}
