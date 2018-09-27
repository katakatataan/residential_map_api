package dto

type PrefCity struct {
	Id       int64   `db:"id" json:"id"`
	CityName string  `db:"city_name" json:"city_name"`
	CityFuri string  `db:"city_furi" json:"city_furi"`
	PrefId   int64   `db:"pref_id" json:"pref_id"`
	Lat      float64 `db:"city_lat" json:"city_lat"`
	Lng      float64 `db:"city_lng" json:"city_lng"`
	PrefName string  `db:"pref_name" json:"pref_name"`
	PrefFuri string  `db:"pref_furi" json:"pref_furi"`
}

type PrefCities []PrefCity
