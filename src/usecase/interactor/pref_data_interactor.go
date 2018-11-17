package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"

	"github.com/k0kubun/pp"
)

type PrefDataInteractor struct {
	PrefDataRankingRepository repository.PrefDataRankingRepository
}

func (cdi *PrefDataInteractor) GetPrefDataRanking(prefDataParam *param.PrefDataParamDto) (response.ResStatisticsPrefDatasBuildCountRanking, error) {
	var prefData entity.PrefDatasBuildCountRanking
	prefData, err := cdi.PrefDataRankingRepository.GetMonthlyPrefRankingOfBuildCount(prefDataParam.Begin, prefDataParam.End)
	pp.Println(prefData)
	res := response.ResStatisticsPrefDatasBuildCountRanking{
		Data: prefData,
	}
	if err != nil {
		return response.ResStatisticsPrefDatasBuildCountRanking{}, err
	}
	return res, nil
}
