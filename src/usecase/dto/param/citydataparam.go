package param

import "time"

type CityDataParamDto struct {
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	From time.Time `query:"from" validate:"required"`
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	To     time.Time `query:"to" validate:"required"`
	PrefId int64     `query:"pref_id"`
	CityId int64     `query:"city_id"`
}
