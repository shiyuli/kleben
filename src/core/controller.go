package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	//get user, it was set by the BasicAuth middleware
	user := context.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		context.HTML(http.StatusOK, "/index.tmpl", gin.H {
			"title": "Main website",
		})
		//context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func Login(context *gin.Context) {
	context.HTML(http.StatusOK, "/login.tmpl", gin.H {
		"title": "Main website",
	})
}
