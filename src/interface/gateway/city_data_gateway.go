package gateway

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/dto"

	"github.com/k0kubun/pp"
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

func (cdg *CityDataGateway) FindByCityId(cityDataParam *dto.CityDataParamDto) (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data WHERE city_id = $1 AND build_date >= $2 AND build_date < $3", cityDataParam.CityId, cityDataParam.From, cityDataParam.To)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) GetMonthlyCityRanking(cityDataParam *dto.CityDataParamDto) (dto.CityDatasDto, error) {
	var cityDatas dto.CityDatasDto
	pp.Println(cityDataParam)
	err := cdg.Find(&cityDatas, "SELECT  * ,rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank FROM city_data WHERE pref_id = $1 AND  build_date >= $2 AND build_date < $3  ORDER BY date_trunc('month', build_date)", cityDataParam.PrefId, cityDataParam.From, cityDataParam.To)
	pp.Println(cityDatas)
	if err != nil {
		return dto.CityDatasDto{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) GetMonthlyPrefRanking(cityDataParam *dto.CityDataParamDto) (dto.CityDatasDto, error) {
	var cityDatas dto.CityDatasDto
	pp.Println(cityDataParam)
	err := cdg.Find(&cityDatas, "SELECT pref_name, SUM(built_count) as built_count, pref_id ,rank() over( partition by date_trunc('month', build_date) order by SUM(built_count) desc) as monthly_rank FROM city_data WHERE build_date >= $1 AND build_date < $2 AND pref_name IS NOT NULL group by pref_id, pref_name, date_trunc('month', build_date) ORDER BY date_trunc('month', build_date)", cityDataParam.From, cityDataParam.To)
	pp.Println(cityDatas)
	if err != nil {
		return dto.CityDatasDto{}, err
	}
	return cityDatas, nil
}
