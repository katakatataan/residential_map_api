package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/interactor/param"
	"residential_map_api/src/usecase/interactor/response"
	"residential_map_api/src/usecase/repository"
)

type PrefDataInteractor struct {
	PrefDataRankingRepository repository.PrefDataRankingRepository
	PrefDataRepository        repository.PrefDataRepository
}

func (cdi *PrefDataInteractor) GetPrefDataRankingBuildCount(param *param.GetPrefsRankingBuildCountParam) (response.GetPrefsRankingBuildCountResponse, error) {
	var prefData entity.PrefDatasBuildCountRanking
	prefData, err := cdi.PrefDataRankingRepository.GetMonthlyPrefRankingOfBuildCount(param.Begin, param.End)
	res := response.GetPrefsRankingBuildCountResponse{
		Data: prefData,
	}
	if err != nil {
		return response.GetPrefsRankingBuildCountResponse{}, err
	}
	return res, nil
}

func (cdi *PrefDataInteractor) GetPrefDataByPrefId(param *param.GetPrefsPrefIdParam) (response.GetPrefsPrefIdResponse, error) {
	var prefData entity.PrefDatas
	prefData, err := cdi.PrefDataRepository.FindByPrefId(param.PrefId, param.Begin, param.End)
	if err != nil {
		return response.GetPrefsPrefIdResponse{}, err
	}
	res := response.GetPrefsPrefIdResponse{
		Data: prefData,
	}
	return res, nil
}
