package entity

import null "gopkg.in/guregu/null.v3"

type CityData struct {
	Id                   int         `db:"id" json:"id"`
	BuiltCount           int         `db:"built_count" json:"built_count"`
	TotalSquareMeter     int         `db:"total_square_meter" json:"total_square_meter"`
	Year                 int         `db:"year" json:"year"`
	Month                int         `db:"month" json:"month"`
	ResidentialUseTypeId int         `db:"residential_use_type_id" json:"residential_use_type_id"`
	ResidentialUseType   string      `db:"residential_use_type" json:"residential_use_type"`
	ConstructionTypeId   int         `db:"construction_type_id" json:"construction_type_id"`
	ConstructionType     string      `db:"construction_type" json:"construction_type"`
	CityId               int         `db:"city_id" json:"city_id"`
	BuidTypeId           int         `db:"build_type_id" json:"build_type_id"`
	BuidType             string      `db:"build_type" json:"build_type"`
	ResidentialTypeId    int         `db:"residential_type_id" json:"residential_type_id"`
	ResidentialType      string      `db:"residential_type" json:"residential_type"`
	StructureTypeId      int         `db:"structure_type_id" json:"structure_type_id"`
	StructureType        string      `db:"structure_type" json:"structure_type"`
	PrefId               int         `db:"pref_id" json:"pref_id"`
	CityName             null.String `db:"city_name" json:"city_name"`
	PrefName             null.String `db:"pref_name" json:"pref_name"`
	BuildDate            string      `db:"build_date" json:"build_date"`
}

type CityDatas []CityData
