package config

import (
	"os"
)

var APP_PORT string
var LOG_PATH string

func init() {

	APP_PORT = os.Getenv("APP_PORT")
	LOG_PATH = os.Getenv("LOG_PATH")
}
