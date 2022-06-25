package main

import (
	"io"

	"fbnoi.com/framework/net/http"
)

func main() {
	engine := http.DefaultEngine()
	engine.GET("homepage", "/", Index, MD1, MD2)
	engine.Run("8080")
}

func MD1(ctx *http.Context, next http.HandleFunc) {
	io.WriteString(ctx.ResponseWriter, "before index in MD1\r\n")
	next(ctx)
	io.WriteString(ctx.ResponseWriter, "after index in MD1\r\n")
}

func MD2(ctx *http.Context, next http.HandleFunc) {
	io.WriteString(ctx.ResponseWriter, "before index in MD2\r\n")
	next(ctx)
	io.WriteString(ctx.ResponseWriter, "after index in MD2\r\n")
}

func Index(ctx *http.Context) {
	io.WriteString(ctx.ResponseWriter, "hello world\r\n")
}
