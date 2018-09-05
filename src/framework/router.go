package framework

import (
	"fmt"
	"residential_map_api/src/interface/controller"
	"residential_map_api/src/usecase/interactor"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn, err := sqlx.Connect("postgres", "user=residential-map password=residential-map dbname=residential sslmode=disable")
	if err != nil {
		fmt.Println("connection error")
	}
	prefCityController := controller.NewMstPrefCityController(interactor.NewMstPrefCityInteractor(conn))

	e.GET("/", prefCityController.GetMstPrefCity)

	// Routes

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
