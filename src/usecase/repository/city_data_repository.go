package repository

import (
	"residential_map_api/src/entity"
)

type CityDataRepository interface {
	FindAll() (entity.CityDatas, error)
	FindById(id int) (entity.CityDatas, error)
}
