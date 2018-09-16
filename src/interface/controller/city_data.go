package controller

import (
	"net/http"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"
)

type cityDataController struct {
	Interactor interactor.CityDataInteractor
}

type CityDataController interface {
	GetCityDataByBuildDate(c Context) error
	GetCityData(c Context) error
}

func NewCityDataController(sqlHandler gateway.SqlHandler) CityDataController {
	return &cityDataController{
		Interactor: interactor.CityDataInteractor{
			Repository: &gateway.CityDataGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *cityDataController) GetCityDataByBuildDate(c Context) error {
	result, err := cd.Interactor.FetchCityDatasByBuildDate()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityData(c Context) error {
	result, err := cd.Interactor.FetchAllCityData()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
