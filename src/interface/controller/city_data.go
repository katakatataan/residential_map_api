package controller

import (
	"net/http"
	"residential_map_api/src/entity"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"
	"strconv"
)

type cityDataController struct {
	Interactor interactor.CityDataInteractor
}

type CityDataController interface {
	GetCityDataByBuildDate(c Context) error
	GetCityData(c Context) error
}

func NewCityDataController(sqlHandler gateway.SqlHandler) *cityDataController {
	return &cityDataController{
		Interactor: interactor.CityDataInteractor{
			Repository: &gateway.CityDataGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *cityDataController) GetCityData(c Context) error {
	e := new(entity.CityData)
	if err := c.Bind(e); err != nil {
		return c.JSON(500, err.Error())
	}
	if err := c.Validate(e); err != nil {
		return c.JSON(400, err.Error())
	}
	result, err := cd.Interactor.FetchAllCityData()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityDataById(c Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := cd.Interactor.FetchCityDatasById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, result)
}
