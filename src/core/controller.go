package core

import (
	"../cloud/qiniu"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	//get user, it was set by the BasicAuth middleware
	user := context.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		context.HTML(http.StatusOK, "index.tmpl", gin.H {
			"title": "Main website",
		})
		//context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func Upload(context *gin.Context) {
	url := context.PostForm("url")

	data, err := downloadFile(url)
	if err != nil {
		print(err.Error())
	}

	key := getUID()
	data = encrypt(data)
	qiniu.Upload(key, data)
}
