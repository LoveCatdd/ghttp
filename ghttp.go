package ghttp

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	routers *router
}

var groups []*RouterGroup

// the constructor of ghttp.Engine
func New() *Engine {
	engine := &Engine{routers: newRouter()}
	engine.RouterGroup = &RouterGroup{routers: engine.routers}
	groups = []*RouterGroup{engine.RouterGroup}

	return engine
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

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	router := group.routers
	newGroup := &RouterGroup{
		prefix:  group.prefix + prefix,
		routers: router,
	}
	groups = append(groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.routers.addRoute(method, pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("get", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.routers.addRoute("post", pattern, handler)
}
