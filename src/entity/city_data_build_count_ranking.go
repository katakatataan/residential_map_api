package entity

import (
	null "gopkg.in/guregu/null.v3"
)

type CityDataBuildCountRanking struct {
	Id                   int         `db:"id" json:"id"`
	BuiltCount           int         `db:"built_count" json:"built_count"`
	TotalSquareMeter     int         `db:"total_square_meter" json:"total_square_meter"`
	Year                 int         `db:"year" json:"year"`
	Month                int         `db:"month" json:"month"`
	ResidentialUseTypeId int         `db:"residential_use_type_id" json:"residential_use_type"`
	ConstructionTypeId   int         `db:"construction_type_id" json:"construction_type_id"`
	CityId               int         `db:"city_id" json:"city_id"`
	BuidTypeId           int         `db:"build_type_id" json:"build_type_id"`
	ResidentialTypeId    int         `db:"residential_type_id" json:"residential_type_id"`
	StructureType        int         `db:"structure_type_id" json:"structure_type"`
	PrefId               int         `db:"pref_id" json:"pref_id"`
	CityName             null.String `db:"city_name" json:"city_name"`
	PrefName             null.String `db:"pref_name" json:"pref_name"`
	BuildDate            string      `db:"build_date" json:"build_date"`
	MontylyRank          int         `db:"monthly_rank" json:"monthly_rank"`
}

type CityDatasBuildCountRanking []CityDataBuildCountRanking
