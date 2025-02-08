package admin

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/user"
	logger "github.com/sirupsen/logrus"
)

type PromoteRequest struct {
	Email string    `json:"email"`
	Role  user.Role `json:"role"`
}

func Promote(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var promote PromoteRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&promote); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	logger.Infof("Email: %s", promote.Email)
	logger.Infof("Role: %v", promote.Role)

	key := user.Key(promote.Email)
	exists, err := storage.GetInstance().Exist().Entity(key).NowT()

	if err != nil {
		http.Error(writer, "Could not check if user exists.", http.StatusInternalServerError)

		return
	}

	if !exists {
		http.Error(writer, "No such user.", http.StatusBadRequest)

		return
	}

	var user user.User

	load_err := storage.GetInstance().Load().Entity(&user, key).Now()

	if load_err != nil {
		http.Error(writer, "Could not load user.", http.StatusInternalServerError)

		return
	}

	is_valid_promotion := promote.Role > user.Role

	if !is_valid_promotion {
		http.Error(writer, "Invalid promotion.", http.StatusBadRequest)

		return
	}

	old_role := user.Role

	user.Role = promote.Role

	save_err := storage.GetInstance().Save().Entity(&user).Now()

	if save_err != nil {
		http.Error(writer, "Could not promote user.", http.StatusInternalServerError)

		return
	}

	logger.Infof("Successfully promoted user: %s to: %v from: %v", user.Username, user.Role, old_role)
}
