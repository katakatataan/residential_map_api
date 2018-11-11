package gateway

import (
	"residential_map_api/src/entity"
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

func (cdg *CityDataGateway) FindByPrefId(prefId int, begin string, end string) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT pref_id, pref_name, SUM(built_count) as built_count, SUM(total_square_meter) AS total_square_meter,  date_trunc('month', build_date) as build_date FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 GROUP BY pref_id, pref_name,  build_date ORDER BY  build_date ASC", prefId, begin, end)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) GetMonthlyCityRankingOfBuildCount(prefId int, begin string, end string) (entity.CityDatasBuildCountRanking, error) {
	var cityDatas entity.CityDatasBuildCountRanking
	err := cdg.Find(&cityDatas, "SELECT  * ,rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank FROM city_data WHERE pref_id = $1 AND  build_date >= $2 AND build_date < $3  ORDER BY date_trunc('month', build_date)", prefId, begin, end)
	if err != nil {
		return entity.CityDatasBuildCountRanking{}, err
	}
	return cityDatas, nil
}
