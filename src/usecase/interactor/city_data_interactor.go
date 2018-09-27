package interactor

import (
	"residential_map_api/src/usecase/dto"
	"residential_map_api/src/usecase/dto/param"
	"residential_map_api/src/usecase/repository"
)

type CityDataInteractor struct {
	Repository repository.CityDataRepository
}

func (cdi *CityDataInteractor) FetchAllCityData() (dto.CityDatas, error) {
	var citydata dto.CityDatas
	citydata, err := cdi.Repository.FindAll()
	if err != nil {
		return dto.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) FetchCityDatasById(cityDataParam *param.CityDataParamDto) (dto.CityDatas, error) {
	var citydata dto.CityDatas
	citydata, err := cdi.Repository.FindByCityId(cityDataParam)
	if err != nil {
		return dto.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) FetchCityDatasByPrefId(cityDataParam *param.CityDataParamDto) (dto.CityDatas, error) {
	var citydata dto.CityDatas
	citydata, err := cdi.Repository.FindByPrefId(cityDataParam)
	if err != nil {
		return dto.CityDatas{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) GetCityDataRanking(cityDataParam *param.CityDataParamDto) (dto.CityDatasDto, error) {
	var citydata dto.CityDatasDto
	citydata, err := cdi.Repository.GetMonthlyCityRanking(cityDataParam)
	if err != nil {
		return dto.CityDatasDto{}, err
	}
	return citydata, nil
}

func (cdi *CityDataInteractor) GetPrefDataRanking(cityDataParam *param.CityDataParamDto) (dto.CityDatasDto, error) {
	var citydata dto.CityDatasDto
	citydata, err := cdi.Repository.GetMonthlyPrefRanking(cityDataParam)
	if err != nil {
		return dto.CityDatasDto{}, err
	}
	return citydata, nil
}
