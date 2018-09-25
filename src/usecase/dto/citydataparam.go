package dto

type CityDataParamDto struct {
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	From string `query:"from" validate:"required"`
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	To     string `query:"to" validate:"required"`
	PrefId int64  `query:"pref_id"`
	CityId int64  `query:"city_id"`
}
