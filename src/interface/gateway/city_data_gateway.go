package gateway

import (
	"residential_map_api/src/entity"
)

type CityDataGateway struct {
	SqlHandler
}

// type CityDataGateway interface {
// 	FindByBuildDate() (entity.CityDatas, error)
// 	FindAll() (entity.CityDatas, error)
// }

// func NewCityDataGateway(sqlHandler SqlHandler) CityDataGateway {
// 	return &cityDataGateway{
// 		Handler: sqlHandler,
// 	}
// }

func (cdg *CityDataGateway) FindByBuildDate() (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	err := cdg.Find(&cityDatas, "SELECT * FROM city_data ORDER BY id ASC limit 10")
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
