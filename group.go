package ghttp

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	routers     *router
}
