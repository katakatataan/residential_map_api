package main

import (
	"site_map_api/src/framework"

	_ "github.com/lib/pq"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Prefecture struct {
	// 注意するのはint64で良いのかkernelで確認する
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Furi string `db:"furi"`
}

func main() {
	// Echo instance
	framework.Run(echo.New())
}
