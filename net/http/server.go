package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"fbnoi.com/httprouter"
	"github.com/pkg/errors"
)

var (
	_default_memory  int64 = 2 * 1024 * 1024
	_default_timeout int64 = 10000
)

func DefaultEngine() *Engine {
	return &Engine{
		router: httprouter.NewRouteTree(&httprouter.Config{RedirectFixedPath: true}),
		config: &Config{_default_memory, _default_timeout},
	}
}

type Config struct {
	MaxMemory int64
	TimeOut   int64
}

type Engine struct {
	server *http.Server

	router *httprouter.RouteTree
	config *Config
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

func resolveAddr(port string) string {
	return fmt.Sprintf(":%s", strings.Trim(port, ":"))
}
