package response

import (
	"github.com/paulmach/go.geojson"
)

type GetCitiesGeojsonBuildCountResponse struct {
	geojson.FeatureCollection
}
