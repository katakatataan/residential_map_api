package framework

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ここはduckタイピングいらないからinterfaceを定義しない
type SqlHandler struct {
	Conn *sqlx.DB
}

func NewSqlHandler() SqlHandler {
	conn, err := sqlx.Connect("postgres", "user=residential-map password=residential-map dbname=residential sslmode=disable")
	if err != nil {
		fmt.Println("connection error")
	}
	return SqlHandler{
		Conn: conn,
	}
}

func (h *SqlHandler) Get(res interface{}, query string, args ...interface{}) error {
	err := h.Conn.Get(res, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (h *SqlHandler) Find(res interface{}, query string, args ...interface{}) error {
	err := h.Conn.Select(res, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (h *SqlHandler) Insert(query string, args ...interface{}) (interface{}, error) {
	result, err := h.Conn.Exec(query, args...)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (h *SqlHandler) Update(query string, args ...interface{}) (interface{}, error) {
	result, err := h.Conn.Exec(query, args...)
	if err != nil {
		return result, err
	}
	return result, nil
}
