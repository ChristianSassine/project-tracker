package db

import (
	"database/sql"

	"github.com/krispier/projectManager/internal/log"
)

// TODO : add database setup
type DB struct {
	DB *sql.DB
}

// const (
// 	user     = os.Getenv("POSTGRES_USER")
// 	dbName   = os.Getenv("POSTGRES_DB")
// 	hostAddr = os.Getenv("POSTGRES_ADDR")
// 	password = os.Getenv("POSTGRES_PASSWORD")
// )

func (db *DB) Connect() {
	connStr := "postgresql://superuser:biguser123@/go?sslmode=disable"
	// Connect to database
	var err error
	db.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.ErrorLog.Fatal(err)
	}
	if err := db.DB.Ping(); err != nil {
		log.ErrorLog.Fatal("Connection to the database failed")
	}
	log.PrintInfo("DB connection successful")
}

func (db *DB) Close() {
	db.Close()
}
