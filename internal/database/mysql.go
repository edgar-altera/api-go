package database

import (
	"database/sql"                     
	"fmt"                              
	"github.com/edgar-altera/api-go/internal/config"
	_"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

)

var DB *sql.DB
var err error

func connect() (db *sql.DB, e error) {

	user := config.MYSQL_USER
	pass := config.MYSQL_PASS
	host := fmt.Sprintf("tcp(%s:%s)", config.MYSQL_HOST, config.MYSQL_PORT)
	ndb  := config.MYSQL_DB
	time := config.MYSQL_TIME

	// Format user:pass@host/db

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=%s", user, pass, host, ndb, time))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func init() {
	
	DB, err = connect()
	
	if err != nil {

		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Fatal("Failed Open MySQL DB")

		return
	}

	// defer DB.Close()

	err = DB.Ping()

	if err != nil {
		
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Fatal("Failed Connect MySQL DB")

		return
	}
	
	log.Info("Connection Successfully to MySQL DB")
}