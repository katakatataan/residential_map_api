package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"
)

type PrefDataInteractor struct {
	PrefDataRankingRepository repository.PrefDataRankingRepository
	PrefDataRepository        repository.PrefDataRepository
}

func (cdi *PrefDataInteractor) GetPrefDataRanking(prefDataParam *param.PrefDataParamDto) (response.ResStatisticsPrefDatasBuildCountRanking, error) {
	var prefData entity.PrefDatasBuildCountRanking
	prefData, err := cdi.PrefDataRankingRepository.GetMonthlyPrefRankingOfBuildCount(prefDataParam.Begin, prefDataParam.End)
	res := response.ResStatisticsPrefDatasBuildCountRanking{
		Data: prefData,
	}
	if err != nil {
		return response.ResStatisticsPrefDatasBuildCountRanking{}, err
	}
	return res, nil
}

func (cdi *PrefDataInteractor) FetchPrefDatasByPrefId(cityDataParam *param.PrefDataParamDto) (response.ResStatisticsCityDatasByPrefId, error) {
	var prefData entity.PrefDatas
	prefData, err := cdi.PrefDataRepository.FindByPrefId(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	if err != nil {
		return response.ResStatisticsCityDatasByPrefId{}, err
	}
	res := response.ResStatisticsCityDatasByPrefId{
		Data: prefData,
	}
	return res, nil
}
