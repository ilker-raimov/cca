package competition

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilker-raimov/cca/common/setup"
	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/model_competition"
	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	"github.com/ilker-raimov/cca/common/util/response"
	"github.com/ilker-raimov/cca/primary/jwt"
	logger "github.com/sirupsen/logrus"
)

func Languages(writer http.ResponseWriter, request *http.Request) {
	languages := setup.All()

	response.WriteOrInternal(writer, languages)
}

var list_role_list = []model_user.Role{model_user.COMPETITOR, model_user.ORGANIZER}

func List(writer http.ResponseWriter, request *http.Request) {
	authorization := request.Header.Get("Authorization")

	email, is_authorization_ok := jwt.ParseAndVerify(authorization, list_role_list, writer)

	if !is_authorization_ok {
		return
	}

	is_all_request := request.URL.Query().Has("all")

	if is_all_request {
		var competitions model_competition.Competitions

		competitions_key := model_competition.KeyAll()
		competitions_fallback := model_competition.NewAll()
		competitions_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&competitions, competitions_key, competitions_fallback).Now()

		if competitions_load_or_create_err != nil {
			response.InternalServerError(writer, model_competition.COULD_NOT_LOAD_OR_CREATE_ALL)

			return
		}

		response.WriteOrInternal(writer, competitions.Ids)
	} else {
		var competitions_user model_competition.CompetitionsUser

		competitions_user_key := model_competition.KeyUser(email)
		competitions_user_fallback := model_competition.NewUser(email)
		competitions_user_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&competitions_user, competitions_user_key, competitions_user_fallback).Now()

		if competitions_user_load_or_create_err != nil {
			response.InternalServerError(writer, model_competition.COULD_NOT_LOAD_OR_CREATE_ALL)

			return
		}

		response.WriteOrInternal(writer, competitions_user.Ids)
	}
}

type CreateRequest struct {
	Public           bool           `json:"public"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Language         setup.Language `json:"language"`
	UseOverallTime   bool           `json:"use_overall_time"`
	UseExecutionTime bool           `json:"use_execution_time"`
}

func Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	authorization := request.Header.Get("Authorization")

	email, is_authorization_ok := jwt.ParseAndVerify(authorization, list_role_list, writer)

	if !is_authorization_ok {
		return
	}

	var create CreateRequest

	is_parse_ok := response.ParseOrInternal(writer, request.Body, &create)

	if !is_parse_ok {
		return
	}

	logger.Infof("Email: %s", email)
	logger.Infof("Title: %s", create.Title)
	logger.Infof("Description: %s", create.Description)
	logger.Infof("Language: %s", create.Language.String())

	competition := model_competition.New(create.Title, create.Description, create.Language)

	logger.Infof("Created competition: %v", competition)

	save_err := storage.GetInstance().Save().Entity(competition).Now()

	if save_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_SAVE)

		return
	}

	logger.Info("Saved competition")

	//user
	var competitions_user model_competition.CompetitionsUser

	competitions_user_key := model_competition.KeyUser(email)
	competitions_user_fallback := model_competition.NewUser(email)
	competitions_user_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&competitions_user, competitions_user_key, competitions_user_fallback).Now()

	if competitions_user_load_or_create_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_LOAD_OR_CREATE_USER)

		return
	}

	competitions_user.Ids = append(competitions_user.Ids, competition.Id)

	competitions_user_save_err := storage.GetInstance().Save().Entity(&competitions_user).Now()

	if competitions_user_save_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_SAVE_USER)

		return
	}

	//all
	var competitions model_competition.Competitions

	competitions_key := model_competition.KeyAll()
	comprtitions_fallback := model_competition.NewAll()
	competitions_load_or_create_err := storage.GetInstance().LoadOrCreate().Entity(&competitions, competitions_key, comprtitions_fallback).Now()

	if competitions_load_or_create_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_LOAD_OR_CREATE_ALL)

		return
	}

	competitions.Ids = append(competitions.Ids, competition.Id)

	competitions_save_err := storage.GetInstance().Save().Entity(&competitions).Now()

	if competitions_save_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_SAVE_ALL)

		return
	}

	writer.Write([]byte(competition.Id))
}

func Get(writer http.ResponseWriter, request *http.Request) {
	authorization := request.Header.Get("Authorization")

	_, is_authorization_ok := jwt.ParseAndVerify(authorization, list_role_list, writer)

	if !is_authorization_ok {
		return
	}

	vars := mux.Vars(request)
	competition_id := vars["competition_id"]

	competition_key := model_competition.Key(competition_id)
	competition_exists, competitions_exist_err := storage.GetInstance().Exist().Entity(competition_key).NowT()

	if competitions_exist_err != nil {
		response.InternalServerError(writer, model_competition.COULT_NOT_CHECK)

		return
	}

	if !competition_exists {
		response.BadRequest(writer, model_competition.NO_SUCH_COMPETITION)

		return
	}

	var competition model_competition.Competition

	competition_load_err := storage.GetInstance().Load().Entity(&competition, competition_key).Now()

	if competition_load_err != nil {
		response.InternalServerError(writer, model_competition.COULD_NOT_LOAD)

		return
	}

	response.WriteOrInternal(writer, competition)
}

func Edit(writer http.ResponseWriter, request *http.Request) {

}

func Delete(writer http.ResponseWriter, request *http.Request) {

}
