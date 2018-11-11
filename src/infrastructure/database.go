package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"residential_map_api/src/interface/gateway"

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

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (gateway.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (gateway.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
