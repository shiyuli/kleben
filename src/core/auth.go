package core

import "github.com/gin-gonic/gin"

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func (router *Router) InitAuth() {
	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	router.group = router.engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
}
