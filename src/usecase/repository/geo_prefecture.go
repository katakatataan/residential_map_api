package repository

import "residential_map_api/src/entity"

type GeoPrefectureRepository interface {
	FindByPrefId(pids []int) (entity.FeatureCollection, error)
}
