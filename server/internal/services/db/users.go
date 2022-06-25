package db

import (
	"errors"
	"strings"

	"github.com/ChristianSassine/projectManager/internal/api"

	_ "github.com/lib/pq"
)

var UserDoesntExistError error = errors.New("User doesn't exist")

func (db *DB) AddUser(username string, password string, email string) error {

	username = strings.ToLower(username)
	_, err := db.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
	return err
}

func (db *DB) CheckIfUserExists(username string) (bool, error) {
	username = strings.ToLower(username)
	count := 0

	if err := db.DB.QueryRow("SELECT COUNT(1) FROM users WHERE users.username = $1", username).Scan(&count); err != nil || count != 1 {
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

	return id, nil
}

func (db *DB) GetUserPassword(id int) (string, error) {
	var password string

	if err := db.DB.QueryRow("SELECT password FROM users WHERE users.id = $1 ", id).Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}
