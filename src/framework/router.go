package framework

import (
	"fmt"
	"residential_map_api/src/interface/controller"
	"residential_map_api/src/usecase/interactor"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/sevenNt/echo-pprof"

	_ "github.com/lib/pq"
)

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// 認証をつけル時のデバッグ用に作成
		fmt.Printf("%s\n", reqBody)
	}))

	conn, err := sqlx.Connect("postgres", "user=residential-map password=residential-map dbname=residential sslmode=disable")
	if err != nil {
		fmt.Println("connection error")
	}
	prefCityController := controller.NewMstPrefCityController(interactor.NewMstPrefCityInteractor(conn))

	e.GET("/", prefCityController.GetMstPrefCity)
	echopprof.Wrap(e)

	// Routes

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
