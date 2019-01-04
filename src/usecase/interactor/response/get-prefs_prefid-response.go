package response

import (
	"residential_map_api/src/entity"
)

type GetPrefsPrefIdResponse struct {
	Data entity.PrefDatas `json:"data"`
}
