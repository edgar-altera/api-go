package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var APP_NAME string
var APP_PORT string
var APP_LANG string

var JWT_SEED string
var JWT_TIME_EXPIRATION_UNIT string
var JWT_TIME_EXPIRATION_VALUE int

var MYSQL_HOST string
var MYSQL_PORT string
var MYSQL_DB string
var MYSQL_USER string
var MYSQL_PASS string
var MYSQL_TIME string

var LOG_PATH string
var LOG_MAX_SIZE_IN_MB int
var LOG_MAX_BACKUPS int
var LOG_MAX_AGE int 
var LOG_COMPRESS bool 
var LOG_CONSOLE_PRINT bool
var LOG_JSON_FORMAT bool
var LOG_SLACK_HOOK_URL string
var LOG_SLACK_USERNAME string
var LOG_SLACK_EMOJI string
var LOG_SLACK_LEVEL_THRESHOLD string
var LOG_SLACK_CHANNEL string

func init() {

	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	log.Println("Loaded successfully env.go")
}

func init() {

	APP_NAME = os.Getenv("APP_NAME")
	APP_PORT = os.Getenv("APP_PORT")
	APP_LANG = os.Getenv("APP_LANG")
	
	JWT_SEED = os.Getenv("JWT_SEED")
	JWT_TIME_EXPIRATION_UNIT = os.Getenv("JWT_TIME_EXPIRATION_UNIT")
	JWT_TIME_EXPIRATION_VALUE, _ = strconv.Atoi(os.Getenv("JWT_TIME_EXPIRATION_VALUE"))

	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	MYSQL_DB   = os.Getenv("MYSQL_DB")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASS = os.Getenv("MYSQL_PASS")
	MYSQL_TIME = os.Getenv("MYSQL_TIME")

	LOG_PATH = os.Getenv("LOG_PATH")
	LOG_CONSOLE_PRINT, _ = strconv.ParseBool(os.Getenv("LOG_CONSOLE_PRINT"))
	LOG_MAX_SIZE_IN_MB, _ = strconv.Atoi(os.Getenv("LOG_MAX_SIZE_IN_MB"))
	LOG_MAX_BACKUPS, _ = strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	LOG_MAX_AGE, _ = strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	LOG_COMPRESS, _ = strconv.ParseBool(os.Getenv("LOG_COMPRESS"))
	LOG_JSON_FORMAT, _ = strconv.ParseBool(os.Getenv("LOG_JSON_FORMAT"))
	LOG_SLACK_HOOK_URL = os.Getenv("LOG_SLACK_HOOK_URL")
	LOG_SLACK_USERNAME = os.Getenv("LOG_SLACK_USERNAME")
	LOG_SLACK_EMOJI = os.Getenv("LOG_SLACK_EMOJI")
	LOG_SLACK_LEVEL_THRESHOLD = os.Getenv("LOG_SLACK_LEVEL_THRESHOLD")
	LOG_SLACK_CHANNEL = os.Getenv("LOG_SLACK_CHANNEL")
}