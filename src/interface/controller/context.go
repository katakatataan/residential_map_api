package controller

import "io"

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
	// HTML sends an HTTP response with status code.
	HTML(code int, html string) error

	// HTMLBlob sends an HTTP blob response with status code.
	HTMLBlob(code int, b []byte) error

	// String sends a string response with status code.
	String(code int, s string) error

	// JSONPretty sends a pretty-print JSON with status code.
	JSONPretty(code int, i interface{}, indent string) error

	// JSONBlob sends a JSON blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP sends a JSONP response with status code. It uses `callback` to construct
	// the JSONP payload.
	JSONP(code int, callback string, i interface{}) error

	// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
	// to construct the JSONP payload.
	JSONPBlob(code int, callback string, b []byte) error

	// XML sends an XML response with status code.
	XML(code int, i interface{}) error

	// XMLPretty sends a pretty-print XML with status code.
	XMLPretty(code int, i interface{}, indent string) error

	// XMLBlob sends an XML blob response with status code.
	XMLBlob(code int, b []byte) error

	// Blob sends a blob response with status code and content type.
	Blob(code int, contentType string, b []byte) error

	// Stream sends a streaming response with status code and content type.
	Stream(code int, contentType string, r io.Reader) error

	// File sends a response with the content of the file.
	File(file string) error
}
