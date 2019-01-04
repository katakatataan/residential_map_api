package response

import (
	"residential_map_api/src/entity"
)

type GetMasterPrefResponse struct {
	Data entity.Prefs `json:"data"`
}
