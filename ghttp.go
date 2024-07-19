package ghttp

import (
	"fmt"
	"net/http"
	"strings"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routers map[string]HandleFunc
}

// the constructor of ghttp.Engine
func New() *Engine {
	return &Engine{
		routers: make(map[string]HandleFunc),
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	builder := strings.Builder{}
	builder.WriteString(req.Method)
	builder.WriteString("-")
	builder.WriteString(req.URL.Path)
	key := builder.String()
	if handle, ok := engine.routers[key]; ok {
		handle(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (engine *Engine) addRoute(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	engine.routers[key] = handle
}

func (engine *Engine) Get(pattern string, handle HandleFunc) {
	engine.addRoute("get", pattern, handle)
}

func (engine *Engine) Post(pattern string, handle HandleFunc) {
	engine.addRoute("post", pattern, handle)
}

func (engine *Engine) Run(addr string, engineSelf http.Handler) (err error) {
	if engineSelf != nil {
		return http.ListenAndServe(addr, engineSelf)
	}
	return http.ListenAndServe(addr, engine)
}
