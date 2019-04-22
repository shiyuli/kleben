package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	controller := new(Controller)
	router := new(Router)

	engine := gin.Default()
	//engine.LoadHTMLGlob("./templates")
	engine.LoadHTMLFiles("./templates/index.tmpl")
	//engine.Static("/", "./templates")

	router.set(engine, controller)

	engine.Run("localhost:8089")
}

type Router struct {

}

func (router *Router) set(engine *gin.Engine, controller *Controller) {
	engine.GET("/", controller.index)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
}

type Controller struct {

}

func (controller *Controller) index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H {
		"title": "Main website",
	})
}
