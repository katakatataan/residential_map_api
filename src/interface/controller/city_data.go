package controller

import (
	"net/http"
	"residential_map_api/src/usecase/interactor"
)

type cdcontroller struct {
	interactor interactor.CityDataInteractor
}

type CityDataController interface {
	GetCityData(c Context) error
}

func NewCityDataController(it interactor.CityDataInteractor) CityDataController {
	return &cdcontroller{
		interactor: it,
	}
}

func (cd *cdcontroller) GetCityData(c Context) error {
	//ここで実際の取得処理を書く
	result, err := cd.interactor.FetchCityDatasByBuildDate()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
