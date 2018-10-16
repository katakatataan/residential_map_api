package entity

import (
	"github.com/jmoiron/sqlx/types"
)

type MstPrefectureGeojson struct {
	PrefId  int            `db:"pref_id" json:"pref_id"`
	Json    types.JSONText `db:"json" json:"json"`
	PubDate string         `db:"pub_date" json:"pub_date"`
	Weight  float64        `db:"weight" json:"weight"`
}
