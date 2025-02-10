package model_task

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/ilker-raimov/cca/common/util/set"
)

type Task struct {
	Id            string `json:"id"`
	CompetitionId string `json:"competition_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ExecutionTime int16  `json:"execution_time"`
	SetupCode     string `json:"setup_code"`
	UserCode      string `json:"user_code"`
}

type Tasks struct {
	Competition string
	Ids         []string
}

type TestType int

type Test struct {
	Type   TestType
	Input  string
	Output string
}

type TaskTests struct {
	CompetitionId string
	TaskId        string
	Tests         []Test
}

const (
	EXAMPLE TestType = iota
	TEST
)

const (
	NO_SUCH_TASK                   = "No such task."
	COULD_NOT_CHECK                = "Could not check task."
	COULD_NOT_LOAD                 = "Could not load task."
	COULD_NOT_SAVE                 = "Could not save task."
	COULD_NOT_SAVE_ALL             = "Could not save tasks."
	COULD_NOT_SAVE_TESTS           = "Could not save task tests."
	COULD_NOT_LOAD_OR_CREATE_ALL   = "Could not load or create tasks."
	COULD_NOT_LOAD_OR_CREATE_TESTS = "Could not load or create task tests."
)

var stringToType = map[string]TestType{
	"EXAMPLE": EXAMPLE,
	"TEST":    TEST,
}

func TypeFromString(value string) TestType {
	upper_value := strings.ToUpper(value)
	typ := stringToType[upper_value]

	return typ
}

func (t *Task) Key() string {
	return Key(t.CompetitionId, t.Id)
}

func (t *Tasks) Key() string {
	return KeyAll(t.Competition)
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

func (tt *TaskTests) Key() string {
	return KeyTests(tt.CompetitionId, tt.TaskId)
}

func Key(competition_id string, task_id string) string {
	return fmt.Sprintf("storage.model.competitions.%s.task.%s", competition_id, task_id)
}

func KeyAll(id string) string {
	return fmt.Sprintf("storage.model.competition.%s.tasks", id)
}

func KeyTests(competition_id string, task_id string) string {
	return fmt.Sprintf("storage.model.competitions.%s.task.%s.tests", competition_id, task_id)
}

func New(competition_id string, name string, description string, execution_time int16, setup_code string, user_code string) *Task {
	id := uuid.New().String()

	return new(id, competition_id, name, description, execution_time, setup_code, user_code)
}

func NewAll(competition string) *Tasks {
	return newAll(competition, []string{})
}

func NewTests(competition_id string, task_id string) *TaskTests {
	return newTests(competition_id, task_id, []Test{})
}

func new(id string, competition_id string, name string, description string, execution_time int16, setup_code string, user_code string) *Task {
	return &Task{
		Id:            id,
		CompetitionId: competition_id,
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

func newTests(competition_id string, task_id string, tests []Test) *TaskTests {
	return &TaskTests{
		CompetitionId: competition_id,
		TaskId:        task_id,
		Tests:         tests,
	}
}
