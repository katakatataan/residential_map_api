package response

import (
	"residential_map_api/src/entity"
)

type ResStatisticsCityDatasByPrefId struct {
	Data entity.PrefDatas `json:"data"`
}
