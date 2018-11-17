package gateway

import (
	"fmt"
	"os"
	"residential_map_api/src/entity"

	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
)

type CityDataGateway struct {
	SqlHandler
}

func (cdg *CityDataGateway) FindAll() (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	q := "SELECT * FROM city_data limit 1000"
	err := cdg.Find(&cityDatas, q)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindByCityId(cityId int, begin string, end string) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data WHERE city_id = $1 AND build_date >= $2 AND build_date < $3 ORDER BY city_id ASC, build_date ASC", cityId, begin, end)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindByPrefId(prefId int, begin string, end string) (entity.PrefDatas, error) {
	var prefDatas entity.PrefDatas
	err := cdg.Find(&prefDatas, "SELECT pref_id, pref_name, SUM(built_count) as built_count, SUM(total_square_meter) AS total_square_meter,  date_trunc('month', build_date) as build_date FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 GROUP BY pref_id, pref_name,  build_date ORDER BY  build_date ASC", prefId, begin, end)
	if err != nil {
		return entity.PrefDatas{}, err
	}
	return prefDatas, nil
}

func (cdg *CityDataGateway) CompareCitiesInSamePrefecture(prefId int, begin string, end string) (interface{}, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=127.0.0.1 port=5432 sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME")))
	// rows, err := conn.Queryx("SELECT * FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 ORDER BY city_id ASC, build_date ASC", prefId, begin, end)
	rows, err := conn.Query("SELECT id, year, month, residential_use_type_id, construction_type_id, build_type_id, residential_type_id, structure_type_id, pref_id, pref_name, build_date, city_id, city_name, built_count, total_square_meter FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 ORDER BY city_id ASC, build_date ASC", prefId, begin, end)
	type city struct {
		CityId           int         `db:"city_id" json:"city_id"`
		CityName         null.String `db:"city_name" json:"city_name"`
		BuiltCount       int         `db:"built_count" json:"built_count"`
		TotalSquareMeter int         `db:"total_square_meter" json:"total_square_meter"`
		MonthlyRank      int         `db:"montyly_rank" json:"monthly_rank"`
	}
	type common struct {
		Id                   int         `db:"id" json:"id"`
		Year                 int         `db:"year" json:"year"`
		Month                int         `db:"month" json:"month"`
		ResidentialUseTypeId int         `db:"residential_use_type_id" json:"residential_use_type"`
		ConstructionTypeId   int         `db:"construction_type_id" json:"construction_type_id"`
		BuildTypeId          int         `db:"build_type_id" json:"build_type_id"`
		ResidentialTypeId    int         `db:"residential_type_id" json:"residential_type_id"`
		StructureType        int         `db:"structure_type_id" json:"structure_type"`
		PrefId               int         `db:"pref_id" json:"pref_id"`
		PrefName             null.String `db:"pref_name" json:"pref_name"`
		BuildDate            string      `db:"build_date" json:"build_date"`
	}
	type data struct {
		common
		cities []city `json:"cities"`
	}
	res := data{}
	for rows.Next() {
		common := common{}
		city := city{}
		err = rows.Scan(&common.Id, &common.Year, &common.Month, &common.ResidentialUseTypeId, &common.ConstructionTypeId, &common.BuildTypeId, &common.ResidentialTypeId, &common.StructureType, &common.PrefId, &common.PrefName, &common.BuildDate, &city.CityId, &city.CityName, &city.BuiltCount, &city.TotalSquareMeter)
		res.cities = append(res.cities, city)
		res.common = common
		if err != nil {
			return res, err
		}
	}
	// pp.Println(res)
	return res, nil
}

func (cdg *CityDataGateway) GetMonthlyCityRankingOfBuildCount(prefId int, begin string, end string) (entity.CityDatasBuildCountRanking, error) {
	var cityDatas entity.CityDatasBuildCountRanking
	err := cdg.Find(&cityDatas, "SELECT  * ,rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank FROM city_data WHERE pref_id = $1 AND  build_date >= $2 AND build_date < $3  ORDER BY date_trunc('month', build_date)", prefId, begin, end)
	if err != nil {
		return entity.CityDatasBuildCountRanking{}, err
	}
	return cityDatas, nil
}
