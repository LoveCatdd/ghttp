package ghttp

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	routers *router
}

// the constructor of ghttp.Engine
func New() *Engine {
	return &Engine{
		routers: newRouter(),
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.routers.handle(c)
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.routers.addRoute(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("get", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRoute("post", pattern, handler)
}

func (engine *Engine) Run(addr string, engineSelf http.Handler) (err error) {
	if engineSelf != nil {
		err = http.ListenAndServe(addr, engineSelf)
	} else {
		err = http.ListenAndServe(addr, engine)
	}
	return
}
