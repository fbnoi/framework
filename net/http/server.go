package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"fbnoi.com/httprouter"
	"github.com/pkg/errors"
)

func DefaultEngine() *Engine {
	return &Engine{
		router: httprouter.NewRouteTree(&httprouter.Config{RedirectFixedPath: true}),
	}
}

type Engine struct {
	server *http.Server

	router *httprouter.RouteTree
}

func (e *Engine) Run(port string) (err error) {
	defer func() { log.Println(err) }()
	if e.server == nil {
		e.server = &http.Server{
			Addr:    resolveAddr(port),
			Handler: e.router,
		}
	}

	if err = e.server.ListenAndServe(); err != nil {
		return errors.Wrapf(err, "port: %v", port)
	}
	return nil
}

func (e *Engine) GET(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "GET", path, fn, mds...)
}

func (e *Engine) POST(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "POST", path, fn, mds...)
}

func (e *Engine) HEAD(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "HEAD", path, fn, mds...)
}

func (e *Engine) PUT(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "PUT", path, fn, mds...)
}

func (e *Engine) PATCH(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "PATCH", path, fn, mds...)
}

func (e *Engine) DELETE(name, path string, fn HandleFunc, mds ...MD) *Engine {
	return e.handle(name, "DELETE", path, fn, mds...)
}

func (e *Engine) All(name, path string, fn HandleFunc, mds ...MD) *Engine {
	e.router.All(name, path, func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {
		wrapHandler(fn, mds...).Handle(e.Context(r, w, ps))
	})
	return e
}

func (e *Engine) handle(name, method, path string, fn HandleFunc, mds ...MD) *Engine {
	e.router.Handle(name, method, path, func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {
		wrapHandler(fn, mds...).Handle(e.Context(r, w, ps))
	})

	return e
}

func (e *Engine) Context(r *http.Request, w http.ResponseWriter, ps httprouter.Params) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
		Engine:         e,
		RouteParams:    ps,
	}
}

func resolveAddr(port string) string {
	return fmt.Sprintf(":%s", strings.Trim(port, ":"))
}
