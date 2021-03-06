package gateway

import (
	"residential_map_api/src/entity"

	"github.com/k0kubun/pp"
	null "gopkg.in/guregu/null.v3"
)

type CityDataGateway struct {
	SqlHandler
}

func (cdg *CityDataGateway) FindByCityId(cityId int, begin string, end string) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	pp.Println(cityId)
	// TODO: マスターコード変換
	err := cdg.Find(&cityDatas, `SELECT
			id,
			built_count,
			total_square_meter,
			year,
			month,
			residential_use_type_id,
			construction_type_id,
			city_id,
			build_type_id,
			residential_type_id,
			structure_type_id,
			pref_id,
			city_name,
			pref_name,
			to_char(build_date,'YYYY-MM') as build_date
		FROM
			city_data
		WHERE
		city_id = $1
		AND
			build_date >= $2
		AND
			build_date < $3
		ORDER BY city_id ASC, build_date ASC`, cityId, begin, end)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindByPrefId(pref_id int, begin string, end string) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	// TODO: マスターコード変換
	err := cdg.Find(&cityDatas, `SELECT
			cd.id,
			cd.built_count,
			cd.total_square_meter,
			cd.year,
			cd.month,
			cd.residential_use_type_id,
			cd.construction_type_id,
			cd.city_id,
			cd.build_type_id,
			cd.residential_type_id,
			cd.structure_type_id,
			cd.pref_id,
			cd.city_name,
			cd.pref_name,
			to_char(cd.build_date,'YYYY-MM') as build_date,
			rut.name as residential_use_type,
			ct.name as construction_type,
			COALESCE(bt.name, '') as build_type,
			COALESCE(rt.name, '') as residential_type,
			COALESCE(st.name, '') as structure_type
		FROM city_data as cd
		LEFT JOIN
			mst_residential_use_type as rut
		ON
			rut.id = cd.residential_use_type_id
		LEFT JOIN
			mst_construction_type as ct
		ON
			ct.id = cd.construction_type_id
		LEFT JOIN
			mst_build_type as bt
		ON
			bt.id = cd.build_type_id
		LEFT JOIN
			mst_residential_type as rt
		ON
			rt.id = cd.residential_type_id
		LEFT JOIN
			mst_structure_type as st
		ON
			st.id = cd.structure_type_id
		WHERE
			cd.pref_id = $1
		AND
			cd.build_date >= $2
		AND
			cd.build_date < $3
		ORDER BY
			cd.city_id ASC,
			cd.build_date ASC`, pref_id, begin, end)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindCityRankingBuildCount(prefId int, begin string, end string) (interface{}, error) {
	// TODO: ここ今interface作るの面倒なのであとで直す
	// conn, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_HOST")))
	// // TODO: マスターコード変換
	rows, err := cdg.SqlHandler.Query(`SELECT
			id,
			year,
			month,
			residential_use_type_id,
			construction_type_id, build_type_id,
			residential_type_id,
			structure_type_id,
			pref_id, pref_name,
			to_char(build_date,'YYYY-MM') as build_date,
			city_id,
			city_name,
			built_count,
			total_square_meter,
			rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank
		FROM
			city_data
		WHERE
			pref_id = $1
		AND
			build_date >= $2
		AND
			build_date < $3
		ORDER BY build_date ASC, city_id ASC`, prefId, begin, end)
	type City struct {
		CityId           int         `db:"city_id" json:"city_id"`
		CityName         null.String `db:"city_name" json:"city_name"`
		BuiltCount       int         `db:"built_count" json:"built_count"`
		TotalSquareMeter int         `db:"total_square_meter" json:"total_square_meter"`
		MonthlyRank      int         `db:"monthly_rank" json:"monthly_rank"`
	}
	type Common struct {
		Id                   int         `db:"id" json:"id"`
		Year                 int         `db:"year" json:"year"`
		Month                int         `db:"month" json:"month"`
		ResidentialUseTypeId int         `db:"residential_use_type_id" json:"residential_use_type_id"`
		ConstructionTypeId   int         `db:"construction_type_id" json:"construction_type_id"`
		BuildTypeId          int         `db:"build_type_id" json:"build_type_id"`
		ResidentialTypeId    int         `db:"residential_type_id" json:"residential_type_id"`
		StructureTypeId      int         `db:"structure_type_id" json:"structure_type_id"`
		PrefId               int         `db:"pref_id" json:"pref_id"`
		PrefName             null.String `db:"pref_name" json:"pref_name"`
		BuildDate            string      `db:"build_date" json:"build_date"`
	}
	type data struct {
		Common
		Cities []City `json:"cities"`
	}
	res := data{}
	for rows.Next() {
		common := Common{}
		city := City{}
		err = rows.Scan(&common.Id, &common.Year, &common.Month, &common.ResidentialUseTypeId, &common.ConstructionTypeId, &common.BuildTypeId, &common.ResidentialTypeId, &common.StructureTypeId, &common.PrefId, &common.PrefName, &common.BuildDate, &city.CityId, &city.CityName, &city.BuiltCount, &city.TotalSquareMeter, &city.MonthlyRank)
		res.Cities = append(res.Cities, city)
		res.Common = common
		if err != nil {
			return res, err
		}
	}
	// pp.Println(res)
	return res, nil
}
func (cdg *CityDataGateway) FindByCityIdByTargetPeriod(cityId int, begin string, end string) (interface{}, error) {
	// TODO: ここ今interface作るの面倒なのであとで直す
	// TODO: 市区町村のでゼロ件表示に対応できているのか修正
	// TODO: マスターコード変換
	// conn, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_HOST")))
	rows, err := cdg.SqlHandler.Query(`SELECT
			id,
			year,
			month,
			residential_use_type_id,
			construction_type_id,
			build_type_id,
			residential_type_id,
			structure_type_id,
			pref_id,
			pref_name,
			to_char(build_date,'YYYY-MM') as build_date,
			city_id,
			city_name,
			built_count,
			total_square_meter
		FROM
		  city_data
		WHERE
			city_id = $1
		AND
			build_date >= $2
		AND
			build_date < $3
		ORDER BY city_id ASC, build_date ASC`, cityId, begin, end)
	// TODO: jsonにmarshalする時にpropertyを読み取るため大文字で表記
	type Monthly struct {
		BuiltCount       int    `db:"built_count" json:"built_count"`
		TotalSquareMeter int    `db:"total_square_meter" json:"total_square_meter"`
		BuildDate        string `db:"build_date" json:"build_date"`
		Year             int    `db:"year" json:"year"`
		Month            int    `db:"month" json:"month"`
	}
	type Common struct {
		Id                   int         `db:"id" json:"id"`
		ResidentialUseTypeId int         `db:"residential_use_type_id" json:"residential_use_type_id"`
		ConstructionTypeId   int         `db:"construction_type_id" json:"construction_type_id"`
		CityId               int         `db:"city_id" json:"city_id"`
		BuildTypeId          int         `db:"build_type_id" json:"build_type_id"`
		ResidentialTypeId    int         `db:"residential_type_id" json:"residential_type_id"`
		StructureTypeId      int         `db:"structure_type_id" json:"structure_type_id"`
		PrefId               int         `db:"pref_id" json:"pref_id"`
		PrefName             null.String `db:"pref_name" json:"pref_name"`
		CityName             null.String `db:"city_name" json:"city_name"`
	}
	type data struct {
		Common
		Monthly []Monthly `json:"monthly"`
	}
	res := data{}
	for rows.Next() {
		common := Common{}
		monthly := Monthly{}
		err = rows.Scan(&common.Id, &monthly.Year, &monthly.Month, &common.ResidentialUseTypeId, &common.ConstructionTypeId, &common.BuildTypeId, &common.ResidentialTypeId, &common.StructureTypeId, &common.PrefId, &common.PrefName, &monthly.BuildDate, &common.CityId, &common.CityName, &monthly.BuiltCount, &monthly.TotalSquareMeter)
		res.Monthly = append(res.Monthly, monthly)
		res.Common = common
		if err != nil {
			return res, err
		}
	}
	// pp.Println(res)
	return res, nil
}
func (cdg *CityDataGateway) GetMonthlyCityRankingOfBuildCount(prefId int, begin string, end string) (entity.CityDatasBuildCountRanking, error) {
	var cityDatas entity.CityDatasBuildCountRanking
	// TODO: マスターコード変換
	err := cdg.Find(&cityDatas, `SELECT
			id,
			built_count,
			total_square_meter,
			year,
			month,
			residential_use_type_id,
			construction_type_id,
			city_id,
			build_type_id,
			residential_type_id,
			pref_id,
			city_name,
			pref_name,
			build_date,
			rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank
		FROM
			city_data
		WHERE
			pref_id = $1
		AND
			build_date >= $2
		AND
			build_date < $3
		ORDER BY date_trunc('month', build_date)`, prefId, begin, end)
	if err != nil {
		return entity.CityDatasBuildCountRanking{}, err
	}
	return cityDatas, nil
}
