package controller

import (
	"residential_map_api/src/usecase/interactor"
)

type geoPrefectureController struct {
	Interactor interactor.GeoPrefectureInteractor
}

type GeoPrefectureController interface {
	GeoPlainPrefecture
}

func NewPrefectureController() *GeoPlainPrefectureController {
	return *geoPrefectureController{}
}

func GeoPlainPrefecture(c Context) {

}
