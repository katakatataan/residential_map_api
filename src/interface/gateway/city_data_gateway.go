package gateway

import (
	"residential_map_api/src/entity"
)

type CityDataGateway struct {
	SqlHandler
}

func (cdg *CityDataGateway) FindByBuildDate() (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data ORDER BY id ASC limit 3")
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}

func (cdg *CityDataGateway) FindAll() (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data limit 10")
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}
