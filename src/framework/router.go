package framework

import (
	"fmt"
	"net/http"
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
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// 認証をつけル時のデバッグ用に作成
		fmt.Printf("%s\n", reqBody)
	}))
	// http2の通信ができているのかをデバッグする
	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>
		`
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})

	conn, err := sqlx.Connect("postgres", "user=residential-map password=residential-map dbname=residential sslmode=disable")
	if err != nil {
		fmt.Println("connection error")
	}
	prefCityController := controller.NewMstPrefCityController(interactor.NewMstPrefCityInteractor(conn))

	e.GET("/", prefCityController.GetMstPrefCity)
	echopprof.Wrap(e)

	// Routes

	// Start server
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
