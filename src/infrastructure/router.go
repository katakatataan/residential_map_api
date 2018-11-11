package infrastructure

import (
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

	// 都道府県ごとのマップ情報 => geojson/pref
	// 市区町村ごとのマップ情報 => geojson/city?city_id=xx
	// 都道府県ごとの統計情報（棒グラフ） => statistics/pref
	// 市区町村ごとの統計情報（棒グラフ） => statistics/city?pref_id=xx
	// 都道府県ごとの市区町村情報（チェックボックスの文言） => cities?pref_id=xx
	// 月ごとの住宅着工数 => monthly?pref_id=xx&start_month=yy&end_month=zz
	sqlHandler := NewSqlHandler()
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))
	routeForMaster(e.Group("/master"), sqlHandler)
	routeForGeojson(e.Group("/geojson"), sqlHandler)
	routeForStatistics(e.Group("/statistics"), sqlHandler)

	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":1323"))
}
