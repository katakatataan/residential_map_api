package controller

import (
	"net/http"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"
)

type mstPrefCityController struct {
	Interactor interactor.MstPrefCityInteractor
}

type MstPrefCityController interface {
	GetMstPrefCity(c Context) error
}

func NewMstPrefCityController(sqlHandler gateway.SqlHandler) *mstPrefCityController {
	return &mstPrefCityController{
		Interactor: interactor.MstPrefCityInteractor{
			Repository: &gateway.MstPrefCityGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (mpc *mstPrefCityController) GetMstPrefCity(c Context) error {
	result, err := mpc.Interactor.FetchAllPrefCities()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}