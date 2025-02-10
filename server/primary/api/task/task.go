package task

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	model_task "github.com/ilker-raimov/cca/common/storage/model/task"
	"github.com/ilker-raimov/cca/common/util/response"
	"github.com/ilker-raimov/cca/primary/jwt"
	logger "github.com/sirupsen/logrus"
)

type CreateRequest struct {
	Name          string
	Description   string
	ExecutionTime int16
	SetupCode     string
	UserCode      string
}

func List(writer http.ResponseWriter, request *http.Request) {
	authorization := request.Header.Get("Authorization")

	_, is_authorization_ok := jwt.ParseAndVerify(authorization, model_user.ROLE_COMPETE, writer)

	if !is_authorization_ok {
		return
	}

	vars := mux.Vars(request)
	competition_id := vars["competition_id"]

	var tasks model_task.Tasks

	tasks_key := model_task.KeyAll(competition_id)
	tasks_fallback := model_task.NewAll(competition_id)
	tasks_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&tasks, tasks_key, tasks_fallback).Now()

	if tasks_load_or_create_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_LOAD_OR_CREATE_ALL)

		return
	}

	response.WriteOrInternal(writer, tasks.Ids)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	authorization := request.Header.Get("Authorization")

	_, is_authorization_ok := jwt.ParseAndVerify(authorization, model_user.ROLES_ORGANIZER, writer)

	if !is_authorization_ok {
		return
	}

	var create CreateRequest

	is_parse_ok := response.ParseOrInternal(writer, request.Body, create)

	if !is_parse_ok {
		return
	}

	vars := mux.Vars(request)
	competition_id := vars["competition_id"]

	logger.Infof("Name: %s", create.Name)
	logger.Infof("Execution time: %d", create.ExecutionTime)

	task := model_task.New(competition_id, create.Name, create.Description, create.ExecutionTime, create.SetupCode, create.UserCode)

	logger.Infof("Created task: %v", task)

	save_err := storage.GetInstance().Save().Entity(task).Now()

	if save_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_SAVE)
	}

	var tasks model_task.Tasks

	tasks_key := model_task.KeyAll(competition_id)
	tasks_fallback := model_task.NewAll(competition_id)
	tasks_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&tasks, tasks_key, tasks_fallback).Now()

	if tasks_load_or_create_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_LOAD_OR_CREATE_ALL)

		return
	}

	tasks.Add(task.Id)

	tasks_save_err := storage.GetInstance().Save().Entity(&tasks).Now()

	if tasks_save_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_SAVE_ALL)
	}
}

func Get(writer http.ResponseWriter, request *http.Request) {
	authorization := request.Header.Get("Authorization")

	_, is_authorization_ok := jwt.ParseAndVerify(authorization, model_user.ROLE_COMPETE, writer)

	if !is_authorization_ok {
		return
	}

	vars := mux.Vars(request)
	competition_id := vars["competition_id"]
	task_id := vars["task_id"]

	task_key := model_task.Key(competition_id, task_id)
	task_exists, task_exist_err := storage.GetInstance().Exist().Entity(task_key).NowT()

	if task_exist_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_CHECK)

		return
	}

	if !task_exists {
		response.BadRequest(writer, model_task.NO_SUCH_TASK)

		return
	}

	var task model_task.Task

	load_err := storage.GetInstance().Load().Entity(&task, task_key).Now()

	if load_err != nil {
		response.InternalServerError(writer, model_task.COULD_NOT_LOAD)
	}

	response.WriteOrInternal(writer, task)
}
