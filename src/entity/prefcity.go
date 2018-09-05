package entity

type PrefCity struct {
	Id       int64   `db:"id"`
	CityName string  `db:"city_name"`
	CityFuri string  `db:"city_furi"`
	PrefId   int64   `db:"pref_id"`
	Lat      float64 `db:"city_lat"`
	Lng      float64 `db:"city_lng"`
	PrefName string  `db:"pref_name"`
	PrefFuri string  `db:"pref_furi"`
}

type PrefCities []PrefCity
