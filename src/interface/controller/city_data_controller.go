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
	FindCityRankingBuildCount(c Context) error
	FindCitiesGeojsonWithBuildCount(c Context) error
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
			GeojsonRepository: &gateway.GeoPrefecture{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *cityDataController) GetCityDataByCityId(c Context) error {
	param := new(param.GetCitiesCityIdParam)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.FindByCityId(param)
	return c.JSON(http.StatusOK, result)
}

func (cd *cityDataController) FindCityRankingBuildCount(c Context) error {
	param := new(param.GetCitiesRankingBuildCountParam)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.FindCityRankingBuildCount(param)
	pp.Println(err)
	jsonBytes, err := json.Marshal(result)
	pp.Println(err)
	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (cd *cityDataController) GetCityDataByTargetPeriod(c Context) error {
	param := new(param.GetCitiesCityIdMonthlyParam)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := cd.Interactor.GetCityDataByTargetPeriod(param)
	jsonBytes, err := json.Marshal(result)
	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (gp *cityDataController) FindCitiesGeojsonWithBuildCount(c Context) error {
	param := new(param.GetCitiesGeojsonBuildCountParam)
	err := c.Bind(param)
	// TODO: コンテキストを引きつづエラーハンドリングに修正
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := gp.Interactor.FindCitiesGeojsonWithBuildCountByPrefId(param)
	return c.JSON(http.StatusOK, result)
}

func (gp *cityDataController) FindCitiesGeojson(c Context) error {
	param := new(param.GetCitiesGeojsonParam)
	err := c.Bind(param)
	// TODO: コンテキストを引きつづエラーハンドリングに修正
	if err != nil {
		return c.JSON(400, err)
	}
	err = c.Validate(param)
	if err != nil {
		return c.JSON(400, err)
	}
	result, err := gp.Interactor.FindCitiesGeojsonByPrefId(param)
	return c.JSON(http.StatusOK, result)
}
