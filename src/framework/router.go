package framework

import (
	"fmt"
	"net/http"
	"residential_map_api/src/interface/controller"
	"residential_map_api/src/usecase/interactor"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
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

func getToken(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "matsu" && password == "yama" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Yuki Matsuyama"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func Run(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// 認証をつけル時のデバッグ用に作成
		fmt.Printf("%s\n", reqBody)
	}))
	e.GET("/token", getToken)
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
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	//　ここの処理は切り分ける、一回のリクエストの中ではコネクションを使いまわして処理する。
	conn, err := sqlx.Connect("postgres", "user=residential-map password=residential-map dbname=residential sslmode=disable")
	if err != nil {
		fmt.Println("connection error")
	}
	prefCityController := controller.NewMstPrefCityController(interactor.NewMstPrefCityInteractor(conn))

	e.GET("/", prefCityController.GetMstPrefCity)
	echopprof.Wrap(e)
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
