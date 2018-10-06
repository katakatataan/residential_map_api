package controller

import (
	"fmt"
	"net/http"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/dto/param"
	"residential_map_api/src/usecase/interactor"

	"github.com/k0kubun/pp"
)

type CityDataController interface {
	GetCityDataByCityId(c Context) error
	GetCityDataByPrefId(c Context) error
	GetCityDataRanking(c Context) error
	GetPrefDataRanking(c Context) error
}

type cityDataController struct {
	Interactor interactor.CityDataInteractor
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

func (cd *cityDataController) GetCityDataByCityId(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		fmt.Println("bind error")
		pp.Println(cityDataParam)
		return c.JSON(404, err)
	}
	fmt.Println("before validate")
	err = c.Validate(cityDataParam)
	if err != nil {
		fmt.Println("validate error")
		return c.JSON(404, err)
	}
	result, err := cd.Interactor.FetchCityDatasById(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityDataByPrefId(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	result, err := cd.Interactor.FetchCityDatasByPrefId(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityDataRanking(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	result, err := cd.Interactor.GetCityDataRanking(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetPrefDataRanking(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	result, err := cd.Interactor.GetPrefDataRanking(cityDataParam)
	return c.JSON(http.StatusOK, result)
}
