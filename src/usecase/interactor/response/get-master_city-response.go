package response

import (
	"residential_map_api/src/entity"
)

type GetMasterCityResponse struct {
	Data entity.Cities `json:"data"`
}
