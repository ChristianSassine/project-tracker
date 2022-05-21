package api

type LoginCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
