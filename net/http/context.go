package http

import (
	"net/http"
	"strconv"
	"strings"

	"fbnoi.com/framework/net/http/binding"
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

func (ctx *Context) Post(name string) string {
	return ctx.Request.PostForm.Get(name)
}

func (ctx *Context) PostInt(name string) (int, error) {
	i := ctx.Post(name)

	return strconv.Atoi(i)
}

func (ctx *Context) PostBool(name string) (bool, error) {
	i := ctx.Post(name)

	return strconv.ParseBool(i)
}

func (ctx *Context) PostFloat(name string) (float64, error) {
	i := ctx.Post(name)

	return strconv.ParseFloat(i, 64)
}

func (ctx *Context) PostSlice(name, sep string) []string {
	i := ctx.Post(name)

	return strings.Split(i, sep)
}

func (ctx *Context) Get(name string) string {
	return ctx.Request.URL.Query().Get(name)
}

func (ctx *Context) GetInt(name string) (int, error) {
	i := ctx.Get(name)

	return strconv.Atoi(i)
}

func (ctx *Context) GetBool(name string) (bool, error) {
	i := ctx.Get(name)

	return strconv.ParseBool(i)
}

func (ctx *Context) GetFloat(name string) (float64, error) {
	i := ctx.Get(name)

	return strconv.ParseFloat(i, 64)
}

func (ctx *Context) GetSlice(name, sep string) []string {
	i := ctx.Get(name)

	return strings.Split(i, sep)
}

func (ctx *Context) Bind(obj any) error {
	b := binding.Default(ctx.Request.Method, ctx.Request.Header.Get("Content-Type"))

	return ctx.BindWith(obj, b)
}

func (ctx *Context) BindWith(obj any, b binding.BindingInterface) error {
	return b.Bind(ctx.Request, obj)
}
