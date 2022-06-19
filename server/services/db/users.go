package db

import (
	"BugTracker/api"
	log "BugTracker/utilities"
	"database/sql"
	"errors"
	"strings"

	_ "github.com/lib/pq"
)

var UserDoesntExistError error = errors.New("User doesn't exist")

func (db *DB) AddUser(username string, password string, email string) (int, error) {
	var id int

	username = strings.ToLower(username)
	var parsedEmail = &sql.NullString{}

	// Adding email if not empty
	if email != "" {
		parsedEmail.String = strings.ToLower(email)
		parsedEmail.Valid = true
		parsedEmail.Value()
	}
	err := db.DB.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", username, parsedEmail, password).Scan(&id)
	return id, err
}

func (db *DB) CheckIfUserExists(username string) (bool, error) {
	username = strings.ToLower(username)
	count := 0

	if err := db.DB.QueryRow("SELECT COUNT(1) users WHERE users.username = $1", username).Scan(&count); err != nil || count != 1 {
		return false, err
	}
	return true, nil
}

func (db *DB) getUser(username string) (*api.User, error) {
	user := &api.User{}

	username = strings.ToLower(username)
	if err := db.DB.QueryRow("SELECT * FROM users WHERE users.username = $1", username).Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) GetUserId(username string) (int, error) {
	var id int

	username = strings.ToLower(username)
	if err := db.DB.QueryRow("SELECT id FROM users WHERE users.username = $1", username).Scan(&id); err != nil {
		return 0, err
	}

	log.PrintInfo("User", username, "is in the database")
	return id, nil
}

func (db *DB) GetUserPassword(id int) (string, error) {
	var password string

	if err := db.DB.QueryRow("SELECT password FROM users WHERE users.id = $1 ", id).Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}
