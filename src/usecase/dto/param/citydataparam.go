package param

type CityDataParamDto struct {
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	// From time.Time `query:"from" validate:"required, canbetime"`
	// To     time.Time `query:"to" validate:"required, canbetime"`
	From   string `query:"from" validate:"required, canbetime"`
	To     string `query:"to" validate:"required, canbetime"`
	PrefId int64  `query:"pref_id"`
	CityId int64  `query:"city_id"`
}
