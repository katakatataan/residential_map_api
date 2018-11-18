package controller

import (
	"encoding/json"
	"net/http"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"

	"github.com/k0kubun/pp"
)

type CityDataController interface {
	GetCityDataByCityId(c Context) error
	GetCityDataInSamePrefecture(c Context) error
	GetCityDataRanking(c Context) error
	GetPrefDataRanking(c Context) error
	GetCityDataByTargetPeriod(c Context) error
}

type cityDataController struct {
	Interactor interactor.CityDataInteractor
}

func NewCityDataController(sqlHandler gateway.SqlHandler) *cityDataController {
	return &cityDataController{
		Interactor: interactor.CityDataInteractor{
			CityDataRepository: &gateway.CityDataGateway{
				SqlHandler: sqlHandler,
			},
			CityDataRankingRepository: &gateway.CityDataGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *cityDataController) GetCityDataByCityId(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.FetchCityDatasById(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityDataRanking(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	pp.Println(cityDataParam)
	result, err := cd.Interactor.GetCityDataRanking(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) GetCityDataInSamePrefecture(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.CompareCitiesInSamePrefecture(cityDataParam)
	pp.Println(result)
	jsonBytes, err := json.Marshal(result)
	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (cd *cityDataController) GetCityDataByTargetPeriod(c Context) error {
	cityDataParam := new(param.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.GetCityDataByTargetPeriod(cityDataParam)
	pp.Println(result)
	jsonBytes, err := json.Marshal(result)
	return c.JSONBlob(http.StatusOK, jsonBytes)
}
