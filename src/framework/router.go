package framework

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/geo", getGeoTokyo)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
	echopprof.Wrap(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
