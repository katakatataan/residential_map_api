package repository

import (
	"residential_map_api/src/entity"
)

type PrefDataRankingRepository interface {
	GetMonthlyPrefRankingOfBuildCount(begin string, end string) (entity.PrefDatasBuildCountRanking, error)
}
