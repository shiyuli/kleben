package core

import (
	"github.com/gin-gonic/gin"
)

type HttpMethod int

const (
	GET HttpMethod = 1 << iota
	POST
)

type Router struct {
	engine *gin.Engine
	group  *gin.RouterGroup
}

func New() *Router {
	return &Router{gin.Default(), nil}
}

func (router *Router) LoadAssets() {
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}

	router.engine.SetHTMLTemplate(t)
}

func (router *Router) LoadHTML() {
	router.engine.LoadHTMLFiles("templates/index.tmpl")
}

func (router *Router) SetS(handler func(*gin.Context), httpMethod HttpMethod, isSecure ...bool) {
	router.Set(handler, httpMethod, true)
}

func (router *Router) Set(handler func(*gin.Context), httpMethod HttpMethod, isSecure ...bool) {
	if len(isSecure) == 0 {
		isSecure = []bool{false}
	}

	if isSecure[0] && router.group == nil {
		throw("Router.group have not initialized yet, please call Router.InitAuth first.")
		return
	}

	var handlerName string
	if handlerName = getFunctionName(handler); handlerName == "index" {
		// hide "index" in url
		handlerName = ""
	}

	switch httpMethod {
	case GET:
		if isSecure[0] {
			router.group.GET("/" + handlerName, handler)
		} else {
			router.engine.GET("/" + handlerName, handler)
		}
		break
	case POST:
		if isSecure[0] {
			router.group.POST("/" + handlerName, handler)
		} else {
			router.engine.POST("/" + handlerName, handler)
		}
		break
	default:
		break
	}
}

func (router *Router) Run(addr string) {
	router.engine.Run(addr)
}
