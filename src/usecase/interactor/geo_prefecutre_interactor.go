package interactor

import (
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"

	geojson "github.com/paulmach/go.geojson"
)

type GeoPrefectureInteractor struct {
	Repository repository.GeoPrefectureRepository
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
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	return res, nil
}
