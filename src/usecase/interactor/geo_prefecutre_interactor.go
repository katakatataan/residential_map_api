package interactor

import (
	"residential_map_api/src/usecase/repository"
)

type GeoPrefectureInteractore struct{
	Repository repository.GeoPrefecutreRepository
}

func(gpi *GeoPrefectureInteractore) (pids []int) entity.FeatureCollection {
	repository.GeoPrefecutreRepository.FindByPrefIds(pids)
}