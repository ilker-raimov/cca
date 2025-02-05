package user

import "fmt"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Key() string {
	return Key(u.Email)
}

func Key(email string) string {
	return fmt.Sprintf("storage.model.user.%s", email)
}

func New(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
