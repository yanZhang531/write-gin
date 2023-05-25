package gee_v0_2

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{make(map[string]HandlerFunc)}
}

func (r *Router) addRouter(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	log.Println("router:", key)
	r.handlers[key] = handler
}

func (r *Router) hanlde(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "%s:%s", "404 pages not found", c.Req.URL)
	}
}
