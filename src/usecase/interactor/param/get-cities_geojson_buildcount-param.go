package param

import (
	"strconv"
	"strings"
)

type GetCitiesGeojsonBuildCountParam struct {
	PrefIds []int   `query:"pref_id" validate:"required"`
	Weight  float64 `query:"weight" validate:"required"`
	Begin   string  `query:"begin" validate:"required,can-be-time"`
	End     string  `query:"end" validate:"required,can-be-time"`
}

// TODO SQLインジェクションはほぼありえないけど、本質的な改修を行う
func (geo *GetCitiesGeojsonBuildCountParam) ToCsvPrefIds() string {
	var s []string
	for _, v := range geo.PrefIds {
		s = append(s, strconv.Itoa(v))
	}
	return strings.Join(s, ",")
}
