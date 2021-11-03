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
	).Info("Logging config successfully")
}

func output() {

    if LOG_CONSOLE_PRINT {
        log.SetOutput(io.MultiWriter(&lumberjack.Logger{
            Filename:   LOG_PATH,
            MaxSize:    LOG_MAX_SIZE_IN_MB, 
            MaxBackups: LOG_MAX_BACKUPS,
            MaxAge:     LOG_MAX_AGE, 
            Compress:   LOG_COMPRESS, 
        }, os.Stdout))
    } else {
        log.SetOutput(&lumberjack.Logger{
            Filename:   LOG_PATH,
            MaxSize:    LOG_MAX_SIZE_IN_MB, 
            MaxBackups: LOG_MAX_BACKUPS,
            MaxAge:     LOG_MAX_AGE, 
            Compress:   LOG_COMPRESS, 
        })
    }
}

func format() {

    if LOG_JSON_FORMAT {

        log.SetFormatter(&log.JSONFormatter{})
    }
}