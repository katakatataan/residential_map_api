package gateway

import (
	"residential_map_api/src/entity"
)

type cityDataGateway struct {
	Handler SqlHandler
}

type CityDataGateway interface {
	FindByBuildDate(entity.CityDatas) error
	FindAll(entity.CityDatas) error
}

func NewCityDataGateway(handler SqlHandler) CityDataGateway {
	return &cityDataGateway{
		Handler: handler,
	}
}

func (cdg *cityDataGateway) FindByBuildDate(cityDatas entity.CityDatas) error {
	err := cdg.Handler.Find(&cityDatas, "SELECT * FROM city_data ORDER BY id ASC limit 10")
	if err != nil {
		return err
	}
	return nil
}

func (cdg *cityDataGateway) FindAll(cityDatas entity.CityDatas) error {
	err := cdg.Handler.Find(&cityDatas, "SELECT * FROM city_data limit 10")
	if err != nil {
		return err
	}
	return nil
}
