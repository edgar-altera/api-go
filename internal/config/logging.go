package config

import (
	log "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
    "fmt"
)

func init() {
    log.SetOutput(&lumberjack.Logger{
        Filename:   LOG_PATH,
        MaxSize:    100, // megabytes
        MaxBackups: 50,
        MaxAge:     31, //days
        Compress:   false, // disabled by default
    })

    log.SetFormatter(&log.JSONFormatter{})

	fmt.Printf("Load logging.go path=%s \n", LOG_PATH)

}