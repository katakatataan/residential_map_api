package response

import (
	"residential_map_api/src/entity"
)

type ResMasterPref struct {
	Data entity.Prefs `json:"data"`
}
