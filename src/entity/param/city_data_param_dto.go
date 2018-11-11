package param

type CityDataParamDto struct {
	Begin  string `query:"begin" validate:"required,can-be-time"`
	End    string `query:"end" validate:"required,can-be-time"`
	PrefId int    `query:"pref_id"`
	CityId int    `query:"city_id"`
}
