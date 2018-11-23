package entity

import null "gopkg.in/guregu/null.v3"

type City struct {
	Id     int        `db:"id" json:"id"`
	PrefId int        `db:"pref_id" json:"pref_id"`
	Lat    null.Float `db:"lat" json:"lat"`
	Lng    null.Float `db:"lng" json:"lng"`
	Name   string     `db:"name" json:"name"`
	Furi   string     `db:"furi" json:"furi"`
}
type Cities []City
