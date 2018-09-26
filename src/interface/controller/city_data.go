package controller

import (
	"net/http"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/dto"
	"residential_map_api/src/usecase/interactor"

	"github.com/k0kubun/pp"
)

type CityDataController struct {
	Interactor interactor.CityDataInteractor
}

func NewCityDataController(sqlHandler gateway.SqlHandler) *CityDataController {
	return &CityDataController{
		Interactor: interactor.CityDataInteractor{
			Repository: &gateway.CityDataGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cd *CityDataController) GetCityDataByCityId(c Context) error {
	cityDataParam := new(dto.CityDataParamDto)
	err := c.Bind(cityDataParam)
	if err != nil {
		return c.JSON(404, err)
	}
	err = c.Validate(cityDataParam)
	if err != nil {
		pp.Println(err)
		return c.JSON(404, err)
	}
	result, err := cd.Interactor.FetchCityDatasById(cityDataParam)
	return c.JSON(http.StatusOK, result)
}

func (cd *CityDataController) GetCityDataRanking(c Context) error {
	cityDataParam := new(dto.CityDataParamDto)
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

func (cd *CityDataController) GetPrefDataRanking(c Context) error {
	cityDataParam := new(dto.CityDataParamDto)
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
