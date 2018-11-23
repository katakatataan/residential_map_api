package response

import (
	"residential_map_api/src/entity"
)

type ResMasterCity struct {
	Data entity.Cities `json:"data"`
}
