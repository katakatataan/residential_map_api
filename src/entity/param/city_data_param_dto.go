package param

type CityDataParamDto struct {
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	//{TODO} ここでstringを受けていることろをbindする際にunmarshalする
	Begin string `query:"begin" validate:"required,can-be-time"`
	End   string `query:"end" validate:"required,can-be-time"`
	// From   string    `query:"from" validate:"required"`
	// To     string    `query:"to" validate:"required"`
	PrefId int `query:"pref_id"`
	CityId int `query:"city_id"`
}
