package db

import (
	"BugTracker/api"
	"BugTracker/utilities"
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "postgresql://superuser:biguser123@/go?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		utilities.ErrorLog.Fatal(err)
	}
	utilities.InfoLog.Println("DB connection successful")
	return db
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
