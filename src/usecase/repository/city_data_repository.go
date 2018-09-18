package repository

import (
	"residential_map_api/src/entity"
)

type CityDataRepository interface {
	FindByBuildDate() (entity.CityDatas, error)
	FindAll() (entity.CityDatas, error)
}
