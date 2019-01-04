package param

type GetCitiesBuildCountParam struct {
	Begin  string `query:"begin" validate:"required,can-be-time"`
	End    string `query:"end" validate:"required,can-be-time"`
	PrefId int    `query:"pref_id"`
}
