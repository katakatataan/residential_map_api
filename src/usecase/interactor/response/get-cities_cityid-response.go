package response

import (
	"residential_map_api/src/entity"
)

type GetCitiesCityIdResponse struct {
	Data entity.CityDatas `json:"data"`
}
