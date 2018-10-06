package controller

type Context interface {
	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error
	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error
	// Validate validates provided `i`. It is usually called after `Context#Bind()`.
	// Validator must be registered using `Echo#Validator`.
	Validate(i interface{}) error
	// Param handles request parameter.
	Param(string) string
}
