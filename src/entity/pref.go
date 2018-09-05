package entity

type Prefectures []Prefecture

type Prefecture struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Furi string `db:"furi"`
}
