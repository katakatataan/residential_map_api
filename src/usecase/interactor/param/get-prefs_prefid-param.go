package param

type GetPrefsPrefIdParam struct {
	Begin  string `query:"begin" validate:"required,can-be-time"`
	End    string `query:"end" validate:"required,can-be-time"`
	PrefId int
}
