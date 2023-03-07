package dto

type UserDto struct {
	ID       int    `json:"id"`
	Email         string `json:"email"`
	PasswordHash string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	Gender        string `json:"gender"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
