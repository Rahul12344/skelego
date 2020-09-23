package skelego

import (
	"net/http"

	"github.com/go-chi/chi"
)

//HTTPRouter interface for routing
type HTTPRouter interface {
	Mount(string, http.HandlerFunc)
	Use(...Middleware)
	Get(string, http.HandlerFunc, ...Middleware)
	Post(string, http.HandlerFunc, ...Middleware)
	Put(string, http.HandlerFunc, ...Middleware)
	Delete(string, http.HandlerFunc, ...Middleware)
}

//Middleware any middleware function
type Middleware = func(next http.Handler) http.Handler

type skelegoChiRouter struct {
	c chi.Router
}

//NewRouter creates new router
func NewRouter(c chi.Router) HTTPRouter {
	return &skelegoChiRouter{
		c: c,
	}
}

//Mounts subroute
func (sr *skelegoChiRouter) Mount(path string, handler http.HandlerFunc) {
	newPath := ""
	if path == "" {
		newPath = "/"
	} else {
		newPath = path
	}
	sr.c.Mount(newPath, handler)
}
func (sr *skelegoChiRouter) Use(middleware ...Middleware) {
	sr.c.With(middleware...)
}
func (sr *skelegoChiRouter) Get(path string, handler http.HandlerFunc, middleware ...Middleware) {
	newPath := ""
	if path == "" {
		newPath = "/"
	} else {
		newPath = path
	}
	sr.c.With(middleware...).Get(newPath, handler)
}
func (sr *skelegoChiRouter) Post(path string, handler http.HandlerFunc, middleware ...Middleware) {
	newPath := ""
	if path == "" {
		newPath = "/"
	} else {
		newPath = path
	}
	sr.c.With(middleware...).Post(newPath, handler)
}
func (sr *skelegoChiRouter) Put(path string, handler http.HandlerFunc, middleware ...Middleware) {
	newPath := ""
	if path == "" {
		newPath = "/"
	} else {
		newPath = path
	}
	sr.c.With(middleware...).Put(newPath, handler)
}

func (sr *skelegoChiRouter) Delete(path string, handler http.HandlerFunc, middleware ...Middleware) {
	newPath := ""
	if path == "" {
		newPath = "/"
	} else {
		newPath = path
	}
	sr.c.With(middleware...).Delete(newPath, handler)
}
