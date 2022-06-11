package db

import (
	"BugTracker/api"
	"BugTracker/utilities"
	"database/sql"
	"errors"
	"strings"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func (db *DB) Connect() {
	connStr := "postgresql://superuser:biguser123@/go?sslmode=disable"
	// Connect to database
	var err error
	db.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		utilities.ErrorLog.Fatal(err)
	}
	utilities.InfoLog.Println("DB connection successful")
}

var UserDoesntExistError error = errors.New("User doesn't exist")

func (db *DB) AddUser(username string, password string, email string) (int, error) {
	var id int

	username = strings.ToLower(username)
	var parsedEmail = &sql.NullString{}

	// Transforming email to null if empty
	if email != "" {
		parsedEmail.String = strings.ToLower(email)
		parsedEmail.Valid = true
		parsedEmail.Value()
	}

	row := db.DB.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", username, parsedEmail, password)
	err := row.Scan(&id)

	return id, err
}

func (db *DB) CheckIfUserExists(username string) (bool, error) {
	username = strings.ToLower(username)
	count := 0

	row := db.DB.QueryRow("SELECT COUNT(1) users WHERE users.username = $1", username)
	err := row.Scan(&count)
	if err != nil || count != 1 {
		return false, err
	}
	return true, nil
}

func (db *DB) getUser(username string) (*api.User, error) {
	user := &api.User{}

	username = strings.ToLower(username)
	row := db.DB.QueryRow("SELECT * FROM users WHERE users.username = $1", username)

	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) ValidateUser(username string, password string) (int, bool) {
	var id int

	username = strings.ToLower(username)
	row := db.DB.QueryRow("SELECT id FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	if err := row.Scan(&id); err != nil {
		return 0, false
	}

	utilities.InfoLog.Println("User", username, "is in the database")
	return id, true
}
