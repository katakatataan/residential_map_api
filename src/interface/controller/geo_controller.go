package controller

import (
	"net/http"
	"residential_map_api/src/entity/param"
	"residential_map_api/src/interface/gateway"
	"residential_map_api/src/usecase/interactor"

	"github.com/k0kubun/pp"
)

type geoPrefectureController struct {
	Interactor interactor.GeoPrefectureInteractor
}

type GeoPrefectureController interface {
	GeoPlainPrefecture(c Context) error
}

func NewGeoPrefectureController(sqlHandler gateway.SqlHandler) *geoPrefectureController {
	return &geoPrefectureController{
		Interactor: interactor.GeoPrefectureInteractor{
			Repository: &gateway.GeoPrefecture{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (gp *geoPrefectureController) GeoPlainPrefecture(c Context) error {
	geoPrefParam := new(param.GeoPrefectureDto)
	err := c.Bind(geoPrefParam)
	if err != nil {
		pp.Println(geoPrefParam)
		return c.JSON(400, err)
	}
	err = c.Validate(geoPrefParam)
	if err != nil {
		pp.Println(geoPrefParam)
		return c.JSON(400, err)
	}
	pp.Println(geoPrefParam)
	result, err := gp.Interactor.FindByPrefId(geoPrefParam)
	return c.JSON(http.StatusOK, result)
}
