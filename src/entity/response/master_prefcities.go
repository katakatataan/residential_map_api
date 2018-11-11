package response

import (
	"residential_map_api/src/entity"
)

type ResMasterPrefCities struct {
	Data entity.PrefCities `json:"data"`
}
