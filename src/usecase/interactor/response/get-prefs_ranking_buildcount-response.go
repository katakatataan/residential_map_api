package response

import (
	"residential_map_api/src/entity"
)

type GetPrefsRankingBuildCountResponse struct {
	Data entity.PrefDatasBuildCountRanking `json:"data"`
}
