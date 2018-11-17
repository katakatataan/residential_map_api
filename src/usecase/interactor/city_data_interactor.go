package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"

	"github.com/k0kubun/pp"
)

type CityDataInteractor struct {
	CityDataRepository        repository.CityDataRepository
	CityDataRankingRepository repository.CityDataRankingRepository
}

func (cdi *CityDataInteractor) FetchAllCityData() (response.ResStatisticsCityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.CityDataRepository.FindAll()
	res := response.ResStatisticsCityDatas{
		Data: citydata,
	}
	if err != nil {
		return response.ResStatisticsCityDatas{}, err
	}
	return res, nil
}

func (cdi *CityDataInteractor) FetchCityDatasById(cityDataParam *param.CityDataParamDto) (response.ResStatisticsCityDatas, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.CityDataRepository.FindByCityId(cityDataParam.CityId, cityDataParam.Begin, cityDataParam.End)
	res := response.ResStatisticsCityDatas{
		Data: citydata,
	}
	if err != nil {
		return response.ResStatisticsCityDatas{}, err
	}
	return res, nil
}

func (cdi *CityDataInteractor) FetchCityDatasByPrefId(cityDataParam *param.CityDataParamDto) (response.ResStatisticsCityDatasByPrefId, error) {
	var prefData entity.PrefDatas
	prefData, err := cdi.CityDataRepository.FindByPrefId(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	pp.Println(prefData)
	if err != nil {
		return response.ResStatisticsCityDatasByPrefId{}, err
	}
	res := response.ResStatisticsCityDatasByPrefId{
		Data: prefData,
	}
	return res, nil
}

func (cdi *CityDataInteractor) CompareCitiesInSamePrefecture(cityDataParam *param.CityDataParamDto) (response.ResStatisticsCityDatasBuildCountInSamePrefecture, error) {
	citydata, err := cdi.CityDataRepository.CompareCitiesInSamePrefecture(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	res := response.ResStatisticsCityDatasBuildCountInSamePrefecture{
		Data: citydata,
	}
	// pp.Println(res)
	if err != nil {
		return response.ResStatisticsCityDatasBuildCountInSamePrefecture{}, err
	}
	return res, nil
}

func (cdi *CityDataInteractor) GetCityDataRanking(cityDataParam *param.CityDataParamDto) (response.ResStatisticsCityDatasBuildCountRanking, error) {
	var citydata entity.CityDatasBuildCountRanking
	citydata, err := cdi.CityDataRankingRepository.GetMonthlyCityRankingOfBuildCount(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	res := response.ResStatisticsCityDatasBuildCountRanking{
		Data: citydata,
	}
	if err != nil {
		return response.ResStatisticsCityDatasBuildCountRanking{}, err
	}
	return res, nil
}
