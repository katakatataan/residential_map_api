package main

import (
	"residential_map_api/src/framework"

	_ "github.com/lib/pq"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	// Echo instance
	framework.Run(echo.New())
}
