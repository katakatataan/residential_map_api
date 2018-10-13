package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/usecase/repository"
)

type GeoPrefectureInteractor struct {
	Repository repository.GeoPrefectureRepository
}

func (gpi *GeoPrefectureInteractor) FindByPrefId(geoPrefParam *param.GeoPrefectureDto) ([]entity.MstPrefectureGeojson, error) {
	result, err := gpi.Repository.FindByPrefId(geoPrefParam)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return result, nil
}
