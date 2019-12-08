package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"os"
)

var Router *gin.Engine

func main()  {

	// Режим работы gin - на продакшене делать "ReleaseMode"
	gin.SetMode(gin.DebugMode)

	// Роутер по-умолчанию
	Router = gin.New()

	// Загрузить шаблоны
	Router.LoadHTMLGlob("templates/*")

	// Загрузить статику
	Router.Static("/css", "css")
	Router.Static("/js", "js")
	Router.Static("/backups", "backups")

	// Проинитить роуты
	InitRoutes(Router)

	// Запустить приложение
	err := Router.Run(":8088")
	if err != nil {
		log.Error("Application not started")
		os.Exit(1)
	}

}
