package user

import (
	"fmt"

	"github.com/ilker-raimov/cca/common/environment"
)

type Role int

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

const (
	COMPETITOR Role = iota
	ORGANIZER
	ADMINISTRATOR
)

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
		Role:     COMPETITOR,
	}
}

func Admin() *User {
	username := environment.GetOrDefault("admin.username", "admin")
	email := environment.GetOrDefault("admin.email", "admin@test.bg")
	password := environment.GetOrDefault("admin.password", "Test123!")

	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     ADMINISTRATOR,
	}
}
