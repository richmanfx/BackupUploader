package main

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.Handle("GET", "/backups", IndexPageShow)
	router.Handle("POST", "/backups/uploading", Uploading)

}