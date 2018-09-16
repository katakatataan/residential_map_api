package entity

import (
	"database/sql"
)

type CityData struct {
	Id                   int64          `db:"id"`
	BuiltCount           int64          `db:"built_count"`
	TotalSquareMeter     int64          `db:"total_square_meter"`
	Year                 int64          `db:"year"`
	Month                int64          `db:"month"`
	ResidentialUseTypeId int64          `db:"residential_use_type_id"`
	ConstructionTypeId   int64          `db:"construction_type_id"`
	CityId               int64          `db:"city_id"`
	BuidTypeId           int64          `db:"build_type_id"`
	ResidentialTypeId    int64          `db:"residential_type_id"`
	StructureType        int64          `db:"structure_type"`
	PrefId               int64          `db:"pref_id"`
	CityName             sql.NullString `db:"city_name"`
	PrefName             sql.NullString `db:"pref_name"`
	BuildDate            string         `db:"build_date"`
}

type CityDatas []CityData
