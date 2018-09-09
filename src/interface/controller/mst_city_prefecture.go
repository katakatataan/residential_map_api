package controller

import (
	"net/http"
	"residential_map_api/src/usecase/interactor"
)

type controller struct {
	interactor interactor.MstPrefCityInteractor
}

type MstPrefCityController interface {
	GetMstPrefCity(c Context) error
}

func NewMstPrefCityController(it interactor.MstPrefCityInteractor) MstPrefCityController {
	return &controller{
		interactor: it,
	}
}

func (mpc *controller) GetMstPrefCity(c Context) error {
	//ここで実際の取得処理を書く
	result, err := mpc.interactor.FetchAllPrefCities()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
