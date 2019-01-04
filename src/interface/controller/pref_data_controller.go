package controller

import (
	"net/http"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"

	"github.com/k0kubun/pp"
)

type PrefDataController interface {
	GetPrefDataRanking(c Context) error
	GetCityDataByPrefId(c Context) error
}

type prefDataController struct {
	Interactor interactor.PrefDataInteractor
}

func NewPrefDataController(sqlHandler gateway.SqlHandler) *prefDataController {
	return &prefDataController{
		Interactor: interactor.PrefDataInteractor{
			PrefDataRankingRepository: &gateway.PrefDataRankingGateway{
				SqlHandler: sqlHandler,
			},
			PrefDataRepository: &gateway.PrefDataGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *prefDataController) GetPrefDataByPrefId(c Context) error {
	// TODO ここは集計したことでprefdataになるのでprefControllerに移行
	param := new(param.GetPrefsPrefIdParam)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	pp.Println(param)
	result, err := cd.Interactor.GetPrefDataByPrefId(param)
	return c.JSON(http.StatusOK, result)
}

func (pd *prefDataController) GetPrefDataRankingBuildCount(c Context) error {
	param := new(param.GetPrefsRankingBuildCountParam)
	// TODO: コンテキストを引き継ぐエラーハンドリングに修正
	err := c.Bind(param)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := pd.Interactor.GetPrefDataRankingBuildCount(param)
	return c.JSON(http.StatusOK, result)
}
