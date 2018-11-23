package entity

type Pref struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Furi string `db:"furi" json:"furi"`
}
type Prefs []Pref
