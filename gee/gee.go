package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// add router
func (e *Engine) addRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET defines the method to add GET request
func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRouter("POST", pattern, handler)
}

// DELETE defines the method to add DELETE request
func (e *Engine) DELETE(pattern string, handler HandleFunc) {
	e.addRouter("DELETE", pattern, handler)
}

// PUT defines the method to add PUT request
func (e *Engine) PUT(pattern string, handler HandleFunc) {
	e.addRouter("PUT", pattern, handler)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {

	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
