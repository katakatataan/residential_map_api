package repository

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/interactor/param"
)

type GeojsonRepository interface {
	FindByPrefId(*param.GetCitiesGeojsonParam) ([]entity.MstPrefectureGeojson, error)
	FindBuildCountByPrefId(prefId int, weight float64, begin string, end string) ([]entity.MstPrefectureGeojson, error)
}
