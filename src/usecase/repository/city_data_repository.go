package repository

import (
	"residential_map_api/src/entity"
)

type CityDataRepository interface {
	FindByCityId(cityId int, begin string, end string) (entity.CityDatas, error)
	FindByPrefId(prefId int, begin string, end string) (entity.CityDatas, error)
	FindCityRankingBuildCount(prefId int, begin string, end string) (interface{}, error)
	FindByCityIdByTargetPeriod(cityId int, begin string, end string) (interface{}, error)
}
