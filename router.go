package main

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.Handle("GET", "/", IndexPageShow)
	router.Handle("POST", "/uploading", Uploading)

}
