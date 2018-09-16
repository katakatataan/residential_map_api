package repository

import (
	"residential_map_api/src/entity"
)

type cdrepo struct {
	handler SqlHandler
}

type CityDataRepository interface {
	FindByBuildDate() (entity.CityDatas, error)
}

func NewCityDataRepository(handler SqlHandler) CityDataRepository {
	return &cdrepo{
		handler: handler,
	}
}

func (r *cdrepo) FindByBuildDate() (entity.CityDatas, error) {
	var cityDatas entity.CityDatas
	// err := r.handler.Find(&cityDatas, "SELECT * FROM city_data WHERE build_date = sysdate() ORDER BY id ASC")
	err := r.handler.Find(&cityDatas, "SELECT * FROM city_data limit 10")
	if err != nil {
		return entity.CityDatas{}, err
	}
	return cityDatas, nil
}
