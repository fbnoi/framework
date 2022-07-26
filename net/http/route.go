package http

import (
	"context"
	"net/http"
	"strings"
	"time"

	"fbnoi.com/handler"
	"fbnoi.com/httprouter"
)

func (e *Engine) GET(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "GET", path, fn, mds...)
}

func (e *Engine) POST(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "POST", path, fn, mds...)
}

func (e *Engine) HEAD(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "HEAD", path, fn, mds...)
}

func (e *Engine) PUT(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "PUT", path, fn, mds...)
}

func (e *Engine) PATCH(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "PATCH", path, fn, mds...)
}

func (e *Engine) DELETE(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	return e.Handle(name, "DELETE", path, fn, mds...)
}

func (e *Engine) All(name, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	h := wrapHandler(fn, mds...)
	e.router.All(name, path, func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {
		e.handle(r, w, ps, h)
	})

	return e
}

func (e *Engine) Handle(name, method, path string, fn func(*Context), mds ...func(*Context, func(*Context))) *Engine {
	h := wrapHandler(fn, mds...)
	e.router.Handle(name, method, path, func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {
		e.handle(r, w, ps, h)
	})

	return e
}

func wrapHandler(fn func(*Context), mds ...func(*Context, func(*Context))) *handler.Handler[*Context] {
	return handler.New[*Context]().Then(mds...).Final(fn)
}

func (e *Engine) handle(r *http.Request, w http.ResponseWriter, ps httprouter.Params, h *handler.Handler[*Context]) {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		r.ParseMultipartForm(_default_memory)
	} else {
		r.ParseForm()
	}

	ct := time.Duration(e.config.TimeOut)

	if t := timeout(r); t < ct && t > 0 {
		ct = t
	}

	var cancel func()
	ctx := &Context{
		Request:        r,
		ResponseWriter: w,
		Engine:         e,
		RouteParams:    ps,
	}

	if ct > 0 {
		ctx.Context, cancel = context.WithTimeout(context.Background(), ct)
	} else {
		ctx.Context, cancel = context.WithCancel(context.Background())
	}

	defer cancel()

	h.Handle(ctx)

}
