package repository

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/dto"
)

type CityDataRepository interface {
	FindAll() (entity.CityDatas, error)
	FindByCityId(*dto.CityDataParamDto) (entity.CityDatas, error)
	GetMonthlyRanking(*dto.CityDataParamDto) (dto.CityDatasDto, error)
}
