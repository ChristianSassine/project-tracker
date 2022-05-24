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

func (db *DB) AddUser(username string, password string, email string) error {
	username = strings.ToLower(username)
	var parsedEmail = &sql.NullString{}
	if email != "" {
		parsedEmail.String = strings.ToLower(email)
		parsedEmail.Valid = true
		parsedEmail.Value()
	}
	_, err := db.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, parsedEmail, password)
	return err
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

func (db *DB) ValidateUser(username string, password string) bool {
	user := &api.User{}

	username = strings.ToLower(username)
	row := db.DB.QueryRow("SELECT * FROM users WHERE users.username = $1 AND users.password = $2", username, password)

	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
		return false
	}

	utilities.InfoLog.Println("User", username, "is in the database")
	return true
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
