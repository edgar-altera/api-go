package config

import (
	log "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
    "io"
    "os"
)

func init() {
    
    output()
    
    format()

    log.WithFields(
		log.Fields{
			"path": LOG_PATH,
		},
	).Info("Loaded successfully logging.go")
}

func output() {

    if LOG_CONSOLE_PRINT {
        log.SetOutput(io.MultiWriter(&lumberjack.Logger{
            Filename:   LOG_PATH,
            MaxSize:    100, // MB
            MaxBackups: 50,
            MaxAge:     31, 
            Compress:   false, 
        }, os.Stdout))
    } else {
        log.SetOutput(&lumberjack.Logger{
            Filename:   LOG_PATH,
            MaxSize:    100, 
            MaxBackups: 50,
            MaxAge:     31, 
            Compress:   false, 
        })
    }
}

func format() {

    if LOG_JSON_FORMAT {

        log.SetFormatter(&log.JSONFormatter{})
    }
}