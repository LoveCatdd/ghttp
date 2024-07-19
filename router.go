package ghttp

import (
	"net/http"

	"github.com/LoveCatdd/ghttp/utils"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handle HandlerFunc) {
	key := utils.Concat(method, pattern, "-")
	r.handlers[key] = handle
}

func (r *router) handle(c *Context) {
	key := utils.Concat(utils.ToLower(c.Method), utils.ToLower(c.Req.URL.Path), "-")
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
