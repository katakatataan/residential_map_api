package infrastructure

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func routeForAuthRequired(g *echo.Group) {
	// Configure middleware with the custom claims type
	//
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	g.Use(middleware.JWTWithConfig(config))
	g.GET("", restricted)
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
