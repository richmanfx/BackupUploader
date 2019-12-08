package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPageShow(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Backups",
		},
	)

}

func Uploading(context *gin.Context) {

	SetLogFormat()

	// Данные из формы
	//fileHeader, err := context.FormFile("file_name")
	_, fileHeader, err := context.Request.FormFile("file_name")

	if err != nil {

		log.Debugf("Ошибка загрузки файла в функции 'Uploading': %s", err)

		context.HTML(http.StatusOK, "message.html",
			gin.H{
				"title":    "Ошибка",
				"message1": "",
				"message2": fmt.Sprintln("Не указан загружаемый файл"),
				"message3": fmt.Sprintf("%s: ", err),
			},
		)

	} else {

		err = context.SaveUploadedFile(fileHeader, "backups/"+fileHeader.Filename)
		if err != nil {
			log.Errorf("Ошибка при загрузке файла '%s': %s", fileHeader.Filename, err)
		}
		log.Debugf("File uploaded: '%s'", fileHeader.Filename)

		context.HTML(http.StatusOK, "message.html",
			gin.H{
				"title":    "Информация",
				"message1": fmt.Sprintf("Файл '%s' успешно загружен", fileHeader.Filename),
				"message2": "",
				"message3": "",
			},
		)

	}

}
