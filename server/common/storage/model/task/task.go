package model_task

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilker-raimov/cca/common/util/set"
)

type Task struct {
	Id            string
	Competition   string
	Name          string
	Description   string
	ExecutionTime int16
	SetupCode     string
	UserCode      string
}

type Tasks struct {
	Competition string
	Ids         []string
}

func (t *Tasks) Has(value string) bool {
	return set.Has(t.Ids, value)
}

func (t *Tasks) Add(value string) {
	t.Ids = set.Add(t.Ids, value)
}

func (t *Tasks) Remove(value string) {
	t.Ids = set.Remove(t.Ids, value)
}

const (
	NO_SUCH_TASK                 = "No such task."
	COULD_NOT_CHECK              = "Could not check task."
	COULD_NOT_LOAD               = "Could not load task."
	COULD_NOT_SAVE               = "Could not save task."
	COULD_NOT_SAVE_ALL           = "Could not save tasks."
	COULD_NOT_LOAD_OR_CREATE_ALL = "Could not load or create tasks."
)

func (t *Task) Key() string {
	return Key(t.Competition, t.Id)
}

func (t *Tasks) Key() string {
	return KeyAll(t.Competition)
}

func Key(competition_id string, id string) string {
	return fmt.Sprintf("storage.model.competitions.%s.task.%s", competition_id, id)
}

func KeyAll(id string) string {
	return fmt.Sprintf("storage.model.competition.%s.tasks", id)
}

func New(competition_id string, name string, description string, execution_time int16, setup_code string, user_code string) *Task {
	id := uuid.New().String()

	return new(id, competition_id, name, description, execution_time, setup_code, user_code)
}

func NewAll(competition string) *Tasks {
	return newAll(competition, []string{})
}

func new(id string, competition_id string, name string, description string, execution_time int16, setup_code string, user_code string) *Task {
	return &Task{
		Id:            id,
		Competition:   competition_id,
		Name:          name,
		Description:   description,
		ExecutionTime: execution_time,
		SetupCode:     setup_code,
		UserCode:      user_code,
	}
}

func newAll(competition string, ids []string) *Tasks {
	return &Tasks{
		Competition: competition,
		Ids:         ids,
	}
}
