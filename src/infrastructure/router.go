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
	//　機能でルーティングを分けるとメンテナンス性が非常に悪くなるやはりドメインに従って設計する
	routeForMaster(e.Group("/master"), sqlHandler)
	// routeForGeojson(e.Group("/geojson"), sqlHandler)
	// routeForStatistics(e.Group("/statistics"), sqlHandler)

	//domainに着目した新しいrouting
	routeForCity(e.Group("/cities"), sqlHandler)
	routeForPref(e.Group("/prefs"), sqlHandler)

	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":1323"))
}
