package http

import (
	"net/http"

	"fbnoi.com/framework/handler"
	"fbnoi.com/httprouter"
)

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
		h := handler.New[*Context]()
		for _, md := range mds {
			h.Then(md)
		}
		h.Final(fn).Handle(e.context(r, w, ps))
	})
	return e
}

func (e *Engine) handle(name, method, path string, fn HandleFunc, mds ...MD) *Engine {
	e.router.Handle(name, method, path, func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {
		h := handler.New[*Context]()
		for _, md := range mds {
			h.Then(md)
		}
		h.Final(fn).Handle(e.context(r, w, ps))
	})

	return e
}
