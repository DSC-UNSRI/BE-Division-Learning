package dto

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCred struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
