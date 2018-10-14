package param

type GeoPrefectureDto struct {
	PrefIds []int   `query:"pref_id" validate:"required"`
	Weight  float64 `query:"weight" validate:"required"`
}
