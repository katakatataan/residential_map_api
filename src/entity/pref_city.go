package entity

import null "gopkg.in/guregu/null.v3"

type PrefCity struct {
	Id       int        `db:"id" json:"id"`
	CityName string     `db:"city_name" json:"city_name"`
	CityFuri string     `db:"city_furi" json:"city_furi"`
	PrefId   int        `db:"pref_id" json:"pref_id"`
	Lat      null.Float `db:"city_lat" json:"city_lat"`
	Lng      null.Float `db:"city_lng" json:"city_lng"`
	PrefName string     `db:"pref_name" json:"pref_name"`
	PrefFuri string     `db:"pref_furi" json:"pref_furi"`
}
type PrefCities []PrefCity
