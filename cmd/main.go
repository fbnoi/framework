package main

import (
	"io"
	"time"

	"fbnoi.com/framework/net/http"
)

func main() {
	engine := http.DefaultEngine()
	engine.GET("homepage", "/", Index, MD1, MD2)
	engine.Run("8080")
}

func MD1(ctx *http.Context, next func(*http.Context)) {
	io.WriteString(ctx.ResponseWriter, "before index in MD1")
	io.WriteString(ctx.ResponseWriter, "\r\n")
	time.Sleep(10 * time.Second)
	next(ctx)
	io.WriteString(ctx.ResponseWriter, "after index in MD1")
	io.WriteString(ctx.ResponseWriter, "\r\n")
}

func MD2(ctx *http.Context, next func(*http.Context)) {
	io.WriteString(ctx.ResponseWriter, "before index in MD2")
	io.WriteString(ctx.ResponseWriter, "\r\n")
	next(ctx)
	io.WriteString(ctx.ResponseWriter, "after index in MD2")
	io.WriteString(ctx.ResponseWriter, "\r\n")
}

func Index(ctx *http.Context) {
	io.WriteString(ctx.ResponseWriter, "----------------------")
	io.WriteString(ctx.ResponseWriter, "\r\n")
	if ctx.Err() != nil {
		io.WriteString(ctx.ResponseWriter, ctx.Err().Error())
	} else {
		io.WriteString(ctx.ResponseWriter, "Done")
	}
	io.WriteString(ctx.ResponseWriter, "\r\n")
	io.WriteString(ctx.ResponseWriter, "----------------------")
	io.WriteString(ctx.ResponseWriter, "\r\n")
}
