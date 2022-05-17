package rin

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine // all groups share a Engine instance
}

// New is the constructor of rin.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// Middleware
func (group *RouterGroup) Use(handlers ...HandlerFunc) *RouterGroup {
	group.middlewares = append(group.middlewares, handlers...)
	return group
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	patternwithprefix := group.prefix + pattern
	log.Printf("Route %4s - %s", method, patternwithprefix)
	group.engine.router.addRoute(method, patternwithprefix, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Implementation of ServeHTTP interface(type Handler) (Need for http.ListenAndServe(,engine))
// Once Implemented, all HTTP request would be handled by the function
// 在 Go 语言中，实现了接口方法的 struct 都可以强制转换为接口类型
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			c.handlers = append(c.handlers, group.middlewares...)
		}
	}
	engine.router.handle(c)
}
