package model_competition

import (
	"fmt"

	"github.com/ilker-raimov/cca/common/setup"

	"github.com/google/uuid"
)

type Competition struct {
	Id          string         `json:"id"`
	Public      bool           `json:"public"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Language    setup.Language `json:"language"`
}

type CompetitionsUser struct {
	User string
	Ids  []string
}

type Competitions struct {
	Ids []string
}

const (
	COULT_NOT_CHECK               = "Could not check if competition exists."
	COULT_NOT_CHECK_USER          = "Could not check if user competitions exist."
	COULT_NOT_CHECK_ALL           = "Could not check if all competitions exist."
	COULD_NOT_LOAD                = "Could not load competition."
	COULD_NOT_LOAD_USER           = "Could not load all user competitions."
	COULD_NOT_LOAD_ALL            = "Could not load all competitions."
	COULD_NOT_SAVE                = "Could not save competition."
	COULD_NOT_SAVE_USER           = "Could not save all user competitions."
	COULD_NOT_SAVE_ALL            = "Could not save all competitions."
	COULD_NOT_LOAD_OR_CREATE      = "Could not create or load competition."
	COULD_NOT_LOAD_OR_CREATE_USER = "Could not create or load user competitions."
	COULD_NOT_LOAD_OR_CREATE_ALL  = "Could not create or load all competitions."
	NO_SUCH_COMPETITION           = "No such competition."
)

func (c *Competition) Key() string {
	return Key(c.Id)
}

func (c *CompetitionsUser) Key() string {
	return KeyUser(c.User)
}

func (c *Competitions) Key() string {
	return KeyAll()
}

func Key(id string) string {
	return fmt.Sprintf("storage.model.competition.%s", id)
}

func KeyUser(id string) string {
	return fmt.Sprintf("storage.model.competitions.%s", id)
}

func KeyAll() string {
	return "storage.model.competitions"
}

func New(title string, description string, language setup.Language) *Competition {
	id := uuid.New().String()

	return new(id, title, description, language)
}

func NewUser(email string) *CompetitionsUser {
	return newUser(email, []string{})
}

func NewAll() *Competitions {
	return newAll([]string{})
}

func new(id string, title string, description string, language setup.Language) *Competition {
	return &Competition{
		Id:          id,
		Title:       title,
		Description: description,
		Language:    language,
	}
}

func newUser(email string, ids []string) *CompetitionsUser {
	return &CompetitionsUser{
		User: email,
		Ids:  ids,
	}
}

func newAll(ids []string) *Competitions {
	return &Competitions{
		Ids: ids,
	}
}
