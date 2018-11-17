package repository

import (
	"residential_map_api/src/entity"
)

type PrefDataRepository interface {
	FindByPrefId(prefId int, begin string, end string) (entity.PrefDatas, error)
}
