package repository

import (
	"residential_map_api/src/entity"
)

type CityDataRankingRepository interface {
	GetMonthlyCityRankingOfBuildCount(prefId int, begin string, end string) (entity.CityDatasBuildCountRanking, error)
}
