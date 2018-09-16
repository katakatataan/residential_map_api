package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type cdinteractor struct {
	Repository repository.CityDataRepository
}

type CityDataInteractor interface {
	FetchCityDatasByBuildDate() (entity.CityDatas, error)
}

func NewCityDataInteractor(repo repository.CityDataRepository) CityDataInteractor {
	return &cdinteractor{
		Repository: repo,
	}
}

func (cdi *cdinteractor) FetchCityDatasByBuildDate() (entity.CityDatas, error) {
	citydatas, err := cdi.Repository.FindByBuildDate()
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydatas, nil
}
