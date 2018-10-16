package infrastructure

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// ここはduckタイピングいらないからinterfaceを定義しない
type SqlHandler struct {
	Conn *sqlx.DB
}

func NewSqlHandler() SqlHandler {

	conn, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=127.0.0.1 port=5432 sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		fmt.Println("connection error")
	}
	if conn == nil {
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

func (h *SqlHandler) In(query string, args ...interface{}) (string, []interface{}, error) {
	q, vs, err := sqlx.In(query, args...)
	return q, vs, err
}
