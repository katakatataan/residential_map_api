package infrastructure

import (
	"residential_map_api/src/interface/controller"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/sevenNt/echo-pprof"
	"gopkg.in/go-playground/validator.v9"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type CustomValidator struct {
	validator *validator.Validate
}

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Validator
	v := validator.New()
	v.RegisterValidation("can-be-time", ValidateTimeString)

	e.Validator = &CustomValidator{validator: v}
	// custombindがいい感じに実装できたらコメントアウト解除
	// e.Binder = NewBinder()
	routeForDebug(e)
	routeForAuthRequired(e.Group("/restricted"))

	sqlHandler := NewSqlHandler()
	mstPrefCityController := controller.NewMstPrefCityController(&sqlHandler)
	cityDataController := controller.NewCityDataController(&sqlHandler)

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
	echopprof.Wrap(e)
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":1323"))
}
