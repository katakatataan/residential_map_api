package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type CityDataInteractor struct {
	Repository repository.CityDataRepository
}

func (cdi *CityDataInteractor) FetchCityDatasByBuildDate() (entity.CityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.Repository.FindByBuildDate()
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) FetchAllCityData() (entity.CityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.Repository.FindAll()
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}
