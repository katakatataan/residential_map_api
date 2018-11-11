package controller

import (
	"net/http"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"
)

type PrefDataController interface {
	GetPrefDataRanking(c Context) error
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
		},
	}
}

func (pd *prefDataController) GetPrefDataRanking(c Context) error {
	prefDataParam := new(param.PrefDataParamDto)
	err := c.Bind(prefDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(prefDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := pd.Interactor.GetPrefDataRanking(prefDataParam)
	return c.JSON(http.StatusOK, result)
}
