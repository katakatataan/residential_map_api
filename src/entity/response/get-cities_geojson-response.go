package response

import (
	"github.com/paulmach/go.geojson"
)

type GetCitiesGeojsonResponse struct {
	geojson.FeatureCollection
}
