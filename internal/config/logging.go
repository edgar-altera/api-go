package config

import (
    "io"
    "os"
	log "github.com/sirupsen/logrus"
    ls "github.com/bluele/logrus_slack"
    "gopkg.in/natefinch/lumberjack.v2"
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

    log.AddHook(&ls.SlackHook{
		HookURL:        LOG_SLACK_HOOK_URL,
		AcceptedLevels: ls.LevelThreshold(getSlackLevelThreshold(LOG_SLACK_LEVEL_THRESHOLD)),
		Channel:        LOG_SLACK_CHANNEL,
		IconEmoji:      LOG_SLACK_EMOJI,
		Username:       LOG_SLACK_USERNAME,
	})
}

func format() {

    if LOG_JSON_FORMAT {

        log.SetFormatter(&log.JSONFormatter{})
    }
}

func getSlackLevelThreshold(ls string) log.Level {

    var level log.Level

    switch ls {
        case "trace":
            level = log.TraceLevel
        case "debug":
            level = log.DebugLevel
        case "info":
            level = log.InfoLevel
        case "warn":
            level = log.WarnLevel
        case "error":
            level = log.ErrorLevel
        case "fatal":
            level = log.FatalLevel
        case "panic":
            level = log.PanicLevel
        default:
            level = log.WarnLevel
    }

    return level
}