package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type cityDataInteractor struct {
	Repository repository.CityDataRepository
}

type CityDataInteractor interface {
	FetchCityDatasByBuildDate() (entity.CityDatas, error)
	FetchAllCityData() (entity.CityDatas, error)
}

func NewCityDataInteractor(repo repository.CityDataRepository) CityDataInteractor {
	return &cityDataInteractor{
		Repository: repo,
	}
}

func (cdi *cityDataInteractor) FetchCityDatasByBuildDate() (entity.CityDatas, error) {
	var citydata entity.CityDatas
	err := cdi.Repository.FindByBuildDate(citydata)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *cityDataInteractor) FetchAllCityData() (entity.CityDatas, error) {
	var citydata entity.CityDatas
	err := cdi.Repository.FindAll(citydata)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}
