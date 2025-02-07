package user

import "fmt"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Key() string {
	return Key(u.Email)
}

func Key(email string) string {
	return fmt.Sprintf("storage.model.user.%s", email)
}

func New(username string, email string, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}
