package http

import (
	"net/http"

	"fbnoi.com/httprouter"
)

type HandleFunc func(*Context)
type MD func(*Context, func(*Context))

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Engine         *Engine
	RouteParams    httprouter.Params
}
