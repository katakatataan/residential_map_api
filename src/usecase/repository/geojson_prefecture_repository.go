package repository

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
)

type GeoPrefectureRepository interface {
	FindByPrefId(*param.GeoPrefectureDto) ([]entity.MstPrefectureGeojson, error)
	FindBuildCountByPrefId(prefId int, weight float64, begin string, end string) ([]entity.MstPrefectureGeojson, error)
}
