package gee_v0_1

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{make(map[string]HandlerFunc)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method, path := r.Method, r.URL.Path
	key := method + "-" + path
	if handler, ok := e.routers[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 pages not found", r.URL)
	}
}

func (e *Engine) addRouter(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	e.routers[key] = handler
}

func (e *Engine) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	e.addRouter("GET", path, handler)
}
func (e *Engine) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	e.addRouter("POST", path, handler)
}
func (e *Engine) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	e.addRouter("PUT", path, handler)
}
func (e *Engine) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	e.addRouter("DELETE", path, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
