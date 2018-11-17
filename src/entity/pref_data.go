package entity

import null "gopkg.in/guregu/null.v3"

type PrefData struct {
	BuiltCount       int         `db:"built_count" json:"build_count"`
	PrefId           int         `db:"pref_id" json:"pref_id"`
	PrefName         null.String `db:"pref_name" json:"pref_name"`
	BuildMonth       string      `db:"build_date" json:"build_date"`
	TotalSquareMeter int         `db:"total_square_meter" json:"total_square_meter"`
}

type PrefDatas []PrefData
