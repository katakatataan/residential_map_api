package main

import (
	"residential_map_api/src/infrastructure"

	_ "github.com/lib/pq"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	// Echo instance
	infrastructure.Run(echo.New())
}
