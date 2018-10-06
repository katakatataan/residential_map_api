package gateway

type SqlHandler interface {
	Get(res interface{}, query string, args ...interface{}) error
	Find(res interface{}, query string, args ...interface{}) error
	Insert(query string, args ...interface{}) (interface{}, error)
	Update(query string, args ...interface{}) (interface{}, error)
}
