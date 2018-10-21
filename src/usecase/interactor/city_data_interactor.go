package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/usecase/repository"
)

type CityDataInteractor struct {
	Repository repository.CityDataRepository
}

func (cdi *CityDataInteractor) FetchAllCityData() (entity.CityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.Repository.FindAll()
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) FetchCityDatasById(cityDataParam *param.CityDataParamDto) (entity.CityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.Repository.FindByCityId(cityDataParam.CityId, cityDataParam.Begin, cityDataParam.End)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) FetchCityDatasByPrefId(cityDataParam *param.CityDataParamDto) (entity.CityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.Repository.FindByPrefId(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	if err != nil {
		return entity.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) GetCityDataRanking(cityDataParam *param.CityDataParamDto) (entity.CityDatasDto, error) {
	var citydata entity.CityDatasDto
	citydata, err := cdi.Repository.GetMonthlyCityRanking(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	if err != nil {
		return entity.CityDatasDto{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) GetPrefDataRanking(cityDataParam *param.CityDataParamDto) (entity.CityDatasDto, error) {
	var citydata entity.CityDatasDto
	citydata, err := cdi.Repository.GetMonthlyPrefRanking(cityDataParam.Begin, cityDataParam.End)
	if err != nil {
		return entity.CityDatasDto{}, err
	}
	return citydata, nil
}
