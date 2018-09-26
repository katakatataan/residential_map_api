package dto

import null "gopkg.in/guregu/null.v3"

type CityDataDto struct {
	Id                   int64       `db:"id" json:"id"`
	BuiltCount           int64       `db:"built_count" json:"built_count"`
	TotalSquareMeter     int64       `db:"total_square_meter" json:"total_square_meter"`
	Year                 int64       `db:"year" json:"year"`
	Month                int64       `db:"month" json:"month"`
	ResidentialUseTypeId int64       `db:"residential_use_type_id" json:"residential_use_type"`
	ConstructionTypeId   int64       `db:"construction_type_id" json:"construction_type_id"`
	CityId               int64       `db:"city_id" json:"city_id"`
	BuidTypeId           int64       `db:"build_type_id" json:"build_type_id"`
	ResidentialTypeId    int64       `db:"residential_type_id" json:"residential_type_id"`
	StructureType        int64       `db:"structure_type" json:"structure_type"`
	PrefId               int64       `db:"pref_id" json:"pref_id"`
	CityName             null.String `db:"city_name" json:"city_name"`
	PrefName             null.String `db:"pref_name" json:"pref_name"`
	BuildDate            string      `db:"build_date" json:"build_date"`
	MontylyRank          int64       `db:"monthly_rank" json:"monthly_rank"`
}

type CityDatasDto []CityDataDto
