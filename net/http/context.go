package http

import (
	"net/http"

	"fbnoi.com/framework/handler"
	"fbnoi.com/httprouter"
)

type HandleFunc handler.HandleFunc[*Context]
type MD func(*Context, HandleFunc)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Engine         *Engine
	RouteParams    httprouter.Params
}

func wrapHandler(fn HandleFunc, mds ...MD) *handler.Handler[*Context] {
	h := handler.New[*Context]()
	for _, md := range mds {
		h.Then(wrapMd(md))
	}
	h.Final(func(ctx *Context) {
		fn(ctx)
	})
	return h
}

func wrapMd(md MD) handler.MD[*Context] {
	return func(ctx *Context, hf handler.HandleFunc[*Context]) {
		md(ctx, func(ctx *Context) {
			hf(ctx)
		})
	}
}
