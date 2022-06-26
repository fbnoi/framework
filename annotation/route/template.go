package route

import (
	"fbnoi.com/example/controller"
	"fbnoi.com/framework/net/http"
)

func RegisterRoute(engine *http.Engine) {
	engine.GET("homepage", "/", controller.Index, controller.MD1, controller.MD2)
}
