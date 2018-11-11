package response

import (
	"residential_map_api/src/entity"
)

type ResStatisticsCityDatas struct {
	Data entity.CityDatas `json:"data"`
}
