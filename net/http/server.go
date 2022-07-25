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

func (e *Engine) context(r *http.Request, w http.ResponseWriter, ps httprouter.Params) *Context {
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
