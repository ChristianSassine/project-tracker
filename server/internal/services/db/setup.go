package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ChristianSassine/projectManager/internal/log"
)

type DB struct {
	DB *sql.DB
}

var (
	user     = os.Getenv("POSTGRES_USER")
	hostAddr = os.Getenv("POSTGRES_ADDR")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbName   = os.Getenv("POSTGRES_DB")
)

func (db *DB) Connect() {
	dbConfig := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostAddr, user, password, dbName)

	// Connect to database
	var err error
	db.DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.ErrorLog.Fatal(err)
	}
	// Trying to connect to the database before stopping
	tries := 5
	currentTry := 0
	for {
		if err := db.DB.Ping(); err != nil {
			log.ErrorLog.Println("Connection to the database failed...Retrying")
			if currentTry >= tries {
				log.ErrorLog.Fatal("Connection to the database failed.")
			}
			currentTry++
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	log.PrintInfo("DB connection successful")
}

func (db *DB) Close() {
	db.Close()
}
