package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

type MasterPrefecture interface {
	FindAll() Prefectures, error
}

type Prefectures []Prefecture

type Prefecture struct {
	// 注意するのはint64で良いのかkernelで確認する
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Furi string `db:"furi"`
}

func getAllPref() []Prefecture {
	db, err := sqlx.Connect("postgres", "user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname="+os.Getenv("DB_NAME")+" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	prefectures := []Prefecture{}
	err = db.Select(&prefectures, "SELECT * FROM mst_pref ORDER BY id ASC")
	pp.Println(prefectures)
	if err != nil {
		fmt.Println(err)
		return []Prefecture{}
	}

	return prefectures
}

func getGeoTokyo(c echo.Context) error {
	resp, err := http.Get("https://storage.googleapis.com/analyze-residential.appspot.com/geo_optimize/201801-1.geojson")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	result := map[string]interface{}{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result)
}
