package gateway

type SqlHandler interface {
	Get(res interface{}, query string, args ...interface{}) error
	Find(res interface{}, query string, args ...interface{}) error
	In(query string, args ...interface{}) (string, []interface{}, error)
	Insert(query string, args ...interface{}) (interface{}, error)
	Update(query string, args ...interface{}) (interface{}, error)
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
