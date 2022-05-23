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

func (db *DB) AddUser(username string, email string, password string) error {
	username = strings.ToLower(username)
	_, err := db.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	return err
}

func (db *DB) getUser(username string) (*api.User, error) {
	user := &api.User{}
	users := []api.User{}

	username = strings.ToLower(username)
	rows, err := db.DB.Query("SELECT * FROM users WHERE users.username = ?", username)

	if err != nil {
		return &api.User{}, err
	}

	for rows.Next() {
		if err := rows.Scan(user.Username, user.Email, user.Password); err != nil {
			return &api.User{}, err
		}
		users = append(users, *user)
	}
	if len(users) != 1 {
		err = errors.New("User doesn't exist ")
		return &api.User{}, UserDoesntExistError
	}
	return user, nil
}

func (db *DB) ValidateUser(username string, password string) (bool, error) {
	user := &api.User{}
	users := []api.User{}

	username = strings.ToLower(username)
	// TODO : use QueryRow instead to simplify code
	rows, err := db.DB.Query("SELECT * FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
			return false, err
		}
		users = append(users, *user)
	}
	if len(users) != 1 {
		err = errors.New("User doesn't exist")
		return false, nil
	}
	utilities.InfoLog.Println("User", username, "is in the database")
	return true, nil
}

func PrintQuery(db *sql.DB) {
	var res api.LoginCreds
	var creds []api.LoginCreds

	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()

	if err != nil {
		utilities.ErrorLog.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&res.Username, &res.Password); err != nil {
			utilities.ErrorLog.Fatal(err)
		}
		creds = append(creds, res)
	}

	utilities.InfoLog.Print(creds)
}
