package entity

import null "gopkg.in/guregu/null.v3"

type PrefDataBuildCountRanking struct {
	Id          int         `db:"id" json:"id"`
	BuiltCount  int         `db:"built_count" json:"built_count"`
	PrefId      int         `db:"pref_id" json:"pref_id"`
	PrefName    null.String `db:"pref_name" json:"pref_name"`
	BuildMonth  string      `db:"build_month" json:"build_month"`
	BuildYear   string      `db:"build_year" json:"build_year"`
	MontylyRank int         `db:"monthly_rank" json:"monthly_rank"`
}

type PrefDatasBuildCountRanking []PrefDataBuildCountRanking
