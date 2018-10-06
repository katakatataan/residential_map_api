package gateway

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
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

func (cdg *CityDataGateway) FindByCityId(cityDataParam *param.CityDataParamDto) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data WHERE city_id = $1 AND build_date >= $2 AND build_date < $3", cityDataParam.CityId, cityDataParam.From, cityDataParam.To)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindByPrefId(cityDataParam *param.CityDataParamDto) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT pref_id, pref_name, SUM(built_count) as built_count, SUM(total_square_meter) AS total_square_meter,  date_trunc('month', build_date) as build_date FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 GROUP BY pref_id, pref_name,  build_date ORDER BY  build_date ASC", cityDataParam.CityId, cityDataParam.From, cityDataParam.To)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) GetMonthlyCityRanking(cityDataParam *param.CityDataParamDto) (entity.CityDatasDto, error) {
	var cityDatas entity.CityDatasDto
	err := cdg.Find(&cityDatas, "SELECT  * ,rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank FROM city_data WHERE pref_id = $1 AND  build_date >= $2 AND build_date < $3  ORDER BY date_trunc('month', build_date)", cityDataParam.PrefId, cityDataParam.From, cityDataParam.To)
	if err != nil {
		return entity.CityDatasDto{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) GetMonthlyPrefRanking(cityDataParam *param.CityDataParamDto) (entity.CityDatasDto, error) {
	var cityDatas entity.CityDatasDto
	err := cdg.Find(&cityDatas, "SELECT pref_name, SUM(built_count) as built_count, pref_id ,rank() over( partition by date_trunc('month', build_date) order by SUM(built_count) desc) as monthly_rank FROM city_data WHERE build_date >= $1 AND build_date < $2 AND pref_name IS NOT NULL group by pref_id, pref_name, date_trunc('month', build_date) ORDER BY date_trunc('month', build_date)", cityDataParam.From, cityDataParam.To)
	if err != nil {
		return entity.CityDatasDto{}, err
	}
	return cityDatas, nil
}
