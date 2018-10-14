package entity

type MstPrefectureGeojson struct {
	PrefId  int     `db:"pref_id" json:"pref_id"`
	Json    string  `db:"json" json:"json"`
	PubDate string  `db:"pub_date" json:"pub_date"`
	Weight  float64 `db:"weight" json:"weight"`
}
