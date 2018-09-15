package framework

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/paulmach/go.geojson"
)

func routeForDebug(e *echo.Echo) {
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// 認証をつけル時のデバッグ用に作成
		fmt.Printf("%s\n", reqBody)
	}))
	e.GET("/geo", func(c echo.Context) error {
		res, err := http.Get("https://storage.googleapis.com/analyze-residential.appspot.com/geo_optimize/201801-1.geojson")
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		body, error := ioutil.ReadAll(res.Body)
		if error != nil {
			log.Fatal(error)
		}
		fc, err := geojson.UnmarshalFeatureCollection(body)
		return c.JSON(200, fc)
	})
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
	e.GET("/token", getToken)
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
