package param

type CityDataParamDto struct {
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	From string `query:"from" validate:"required,can-be-time"`
	To   string `query:"to" validate:"required,can-be-time"`
	// From   string    `query:"from" validate:"required"`
	// To     string    `query:"to" validate:"required"`
	PrefId int64 `query:"pref_id"`
	CityId int64 `query:"city_id"`
}
