package repository

import (
	"residential_map_api/src/entity"
)

type CityDataRepository interface {
	FindAll() (entity.CityDatas, error)
	FindByCityId(cityId int, begin string, end string) (entity.CityDatas, error)
	FindByPrefId(prefId int, begin string, end string) (entity.CityDatas, error)
	CompareCitiesInSamePrefecture(prefId int, begin string, end string) (map[string]interface{}, error)
}
