package infrastructure

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/random"
	_ "github.com/lib/pq"
	"github.com/sevenNt/echo-pprof"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// ServerHeader middleware adds a `Server` header to the response.
func RequestId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestId := c.Request().Header.Get("X-Request-Id")
		pp.Println(requestId)
		if requestId == "" {
			requestId = random.String(32)
		}
		c.Response().Header().Set("X-Request-Id", requestId)
		return next(c)
	}
}

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Recover())
	e.Use(RequestId)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","request_id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"request_id":"${header:[X-Request-Id]}",` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))
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
