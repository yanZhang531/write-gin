package gee_v0_3

import (
	"net/http"
)

type H map[string]any

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.addRoute("GET", path, handler)
}
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.addRoute("POST", path, handler)
}
func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.router.addRoute("PUT", path, handler)
}
func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.router.addRoute("DELETE", path, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
