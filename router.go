package ghttp

import (
	"net/http"
	"strings"

	"github.com/LoveCatdd/ghttp/utils"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func (r *router) addRoute(method string, pattern string, handle HandlerFunc) {
	parts := parsePattern(pattern)

	key := utils.Concat(method, utils.ToLower(pattern), "-")
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handle
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}

	c.Params = params
	key := utils.Concat(utils.ToLower(c.Method), utils.ToLower(n.pattern), "-")
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// 解析路由路径
func parsePattern(pattern string) (parts []string) {
	vs := strings.Split(pattern, "/")

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return
}

func (r *router) getRoute(method string, path string) (n *node, params map[string]any) {
	searchParts := parsePattern(path)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n = root.search(searchParts, 0)
	params = make(map[string]any)

	if n != nil {
		parts := parsePattern(n.pattern)
		for idx, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[idx]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[idx:], "/")
			}
		}
		return
	}
	return nil, nil
}
