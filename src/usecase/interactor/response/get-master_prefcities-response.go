package response

import (
	"residential_map_api/src/entity"
)

type GetMasterPrefCitiesResponse struct {
	Data entity.PrefCities `json:"data"`
}
