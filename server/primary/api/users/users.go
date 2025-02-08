package users

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	"github.com/ilker-raimov/cca/primary/jwt"
)

var me_roles = []model_user.Role{model_user.COMPETITOR, model_user.ORGANIZER, model_user.ADMINISTRATOR}

type MeResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Role     model_user.Role `json:"role"`
}

func Me(writer http.ResponseWriter, request *http.Request) {
	authentication := request.Header.Get("Authentication")
	is_authentication_missing := len(authentication) == 0

	if is_authentication_missing {
		http.Error(writer, "Missing Authentication header.", http.StatusBadRequest)

		return
	}

	email, is_authentication_ok := jwt.ParseAndVerify(authentication, me_roles, writer)

	if !is_authentication_ok {
		return
	}

	key := model_user.Key(email)
	user_exists, exist_err := storage.GetInstance().Exist().Entity(key).NowT()

	if exist_err != nil {
		http.Error(writer, model_user.COULD_NOT_CHECK, http.StatusInternalServerError)

		return
	}

	if !user_exists {
		http.Error(writer, model_user.NO_SUCH_USER, http.StatusInternalServerError)

		return
	}

	var user model_user.User

	load_err := storage.GetInstance().Load().Entity(&user, key).Now()

	if load_err != nil {
		http.Error(writer, model_user.COULD_NOT_LOAD, http.StatusInternalServerError)

		return
	}

	response := MeResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	data, marshal_err := json.Marshal(response)

	if marshal_err != nil {
		http.Error(writer, "Could not marshal response.", http.StatusInternalServerError)

		return
	}

	writer.Write(data)
}
