package repository

import (
	"residential_map_api/src/usecase/dto"
	"residential_map_api/src/usecase/dto/param"
)

type CityDataRepository interface {
	FindAll() (dto.CityDatas, error)
	FindByCityId(*param.CityDataParamDto) (dto.CityDatas, error)
	FindByPrefId(*param.CityDataParamDto) (dto.CityDatas, error)
	GetMonthlyCityRanking(*param.CityDataParamDto) (dto.CityDatasDto, error)
	GetMonthlyPrefRanking(*param.CityDataParamDto) (dto.CityDatasDto, error)
}
