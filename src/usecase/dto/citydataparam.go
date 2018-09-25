package dto

type CityDataParamDto struct {
	From   string `query:"from" validate:"required"`
	To     string `query:"to" validate:"required"`
	PrefId int64  `query:"pref_id"`
	CityId int64  `query:"city_id"`
}
