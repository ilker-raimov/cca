package model_user

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

const (
	COULD_NOT_CHECK = "Could not check if user exists."
	NO_SUCH_USER    = "No such user."
	COULD_NOT_LOAD  = "Could not load user."
)

var (
	ROLES_COMPETITOR    = []Role{COMPETITOR}
	ROLES_ORGANIZER     = []Role{ORGANIZER}
	ROLES_ADMINISTRATOR = []Role{ADMINISTRATOR}
	ROLE_COMPETE        = []Role{COMPETITOR, ORGANIZER}
)

func Roles() map[int]string {
	return map[int]string{
		0: "COMPETITOR",
		1: "ORGANIZER",
		2: "ADMINISTRATOR"}
}

func (u *User) Key() string {
	return Key(u.Email)
}

func Key(email string) string {
	return fmt.Sprintf("storage.model.user.%s", email)
}

func New(username string, email string, password string) *User {
	return new(username, email, password, COMPETITOR)
}

func Competitor() *User {
	username := environment.GetOrDefault("competitor.username", "competitor")
	email := environment.GetOrDefault("competitor.email", "competitor@test.bg")
	password := environment.GetOrDefault("competitor.password", "Test123!")

	return new(username, email, password, COMPETITOR)
}

func Organizer() *User {
	username := environment.GetOrDefault("organizer.username", "organizer")
	email := environment.GetOrDefault("organizer.email", "organizer@test.bg")
	password := environment.GetOrDefault("organizer.password", "Test123!")

	return new(username, email, password, ORGANIZER)
}

func Admin() *User {
	username := environment.GetOrDefault("admin.username", "admin")
	email := environment.GetOrDefault("admin.email", "admin@test.bg")
	password := environment.GetOrDefault("admin.password", "Test123!")

	return new(username, email, password, ADMINISTRATOR)
}

func new(username string, email string, password string, role Role) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
}
