package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"

	logger "github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var login LoginRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&login); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	logger.Infof("Email: %s", login.Email)
	logger.Infof("Password: %s", login.Password)

	storage.GetInstance().Load(login.Email)

	logger.Info("Successful login")
}

func Register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var register RegisterRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&register); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	logger.Infof("Name: %s", register.Name)
	logger.Infof("Email: %s", register.Email)
	logger.Infof("Password: %s", register.Password)

	exists, err := storage.GetInstance().Exists(register.Email)

	if err != nil {
		logger.Infof("Could not check if key %s exists due to: %s", register.Email, err.Error())
	}

	if exists {
		msg := "User with this email already exists."

		logger.Info(msg)

		http.Error(writer, msg, http.StatusBadRequest)

		return
	}

	storage.GetInstance().Save(register.Email, []byte(register.Name))

	logger.Info("Successful register")
}
