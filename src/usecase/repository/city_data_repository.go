package repository

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
)

type CityDataRepository interface {
	FindAll() (entity.CityDatas, error)
	FindByCityId(*param.CityDataParamDto) (entity.CityDatas, error)
	FindByPrefId(*param.CityDataParamDto) (entity.CityDatas, error)
	GetMonthlyCityRanking(*param.CityDataParamDto) (entity.CityDatasDto, error)
	GetMonthlyPrefRanking(*param.CityDataParamDto) (entity.CityDatasDto, error)
}
