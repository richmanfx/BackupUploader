package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPageShow(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H {
			"title":        "Backups",
		},
	)

}

func Uploading(context *gin.Context) {

	// Данные из формы
	fileName := context.PostForm("file)name")
	log.Debugf("File name: '%s'", fileName)


}
