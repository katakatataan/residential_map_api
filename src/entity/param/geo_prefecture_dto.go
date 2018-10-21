package param

import (
	"strconv"
	"strings"
)

type GeoPrefectureDto struct {
	PrefIds []int   `query:"pref_id" validate:"required"`
	Weight  float64 `query:"weight" validate:"required"`
}

// TODO SQLインジェクションはほぼありえないけど、本質的な改修を行う
func (geo *GeoPrefectureDto) ToCsvPrefIds() string {
	var s []string
	for _, v := range geo.PrefIds {
		s = append(s, strconv.Itoa(v))
	}
	return strings.Join(s, ",")
}
