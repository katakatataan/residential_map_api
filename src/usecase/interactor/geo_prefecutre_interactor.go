package interactor

import (
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"
	"strconv"

	"github.com/k0kubun/pp"
	geojson "github.com/paulmach/go.geojson"
)

type GeoPrefectureInteractor struct {
	Repository         repository.GeoPrefectureRepository
	CityDataRepository repository.CityDataRepository
}

func (gpi *GeoPrefectureInteractor) FindByPrefId(geoPrefParam *param.GeoPrefectureDto) (response.ResGeojsonFeatureCollection, error) {
	result, err := gpi.Repository.FindByPrefId(geoPrefParam)
	var res response.ResGeojsonFeatureCollection
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			return response.ResGeojsonFeatureCollection{}, err
		}
		for _, v := range fc.Features {
			// TODO: ここのloopのレベルでデータを結合
			pp.Println(v.Properties["N03_007"])
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	return res, nil
}

func (gpi *GeoPrefectureInteractor) FindBuildCountByPrefId(geoPrefParam *param.GeoPrefectureWithPeriodDto) (response.ResGeojsonFeatureCollection, error) {
	// TODO: 後々に複数の都道府県に対応する
	result, err := gpi.Repository.FindBuildCountByPrefId(geoPrefParam.PrefIds[0], geoPrefParam.Weight, geoPrefParam.Begin, geoPrefParam.End)
	var res response.ResGeojsonFeatureCollection
	citydata, _ := gpi.CityDataRepository.FindByPrefId(geoPrefParam.PrefIds[0], geoPrefParam.Begin, geoPrefParam.End)
	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	// FIXME : 処理が無駄に回ってしまうので最適化
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			return response.ResGeojsonFeatureCollection{}, err
		}
		for _, v := range fc.Features {
			// FIXME : 処理が無駄に回ってしまうので最適化
			cityId := v.Properties["N03_007"].(string)
			cityIdInt, _ := strconv.Atoi(cityId)
			for _, d := range citydata {
				if d.CityId == cityIdInt {
					v.Properties["statistics"] = d
					break
				}
			}
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	return res, nil
}
