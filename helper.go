package main

import (
	log "github.com/Sirupsen/logrus"
)

// Установить формат логов и уровень логирования
func SetLogFormat() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	log.SetLevel(log.DebugLevel)
}
