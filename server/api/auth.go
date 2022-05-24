package api

import "database/sql"

type LoginCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Email    sql.NullString `json:"email"`
}
