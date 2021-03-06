package controller

import (
	"net/http"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"
	"residential_map_api/src/usecase/interactor/param"
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
	param := new(param.GetMasterPrefCitiesParam)
	err := c.Bind(param)
	// TOOD: コンテキストを引き継ぐエラーハンドリングに修正
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := mpc.Interactor.FetchAllPrefCities(param)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (mpc *mstPrefCityController) GetMstPref(c Context) error {
	result, err := mpc.Interactor.FetchPref()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (mpc *mstPrefCityController) GetMstCity(c Context) error {
	result, err := mpc.Interactor.FetchCity()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
