package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/user"

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

	key := user.Key(login.Email)
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

	storage.GetInstance().Load().Entity(&user, key).Now()

	match := user.Password == login.Password

	if !match {
		http.Error(writer, "Invalid password.", http.StatusBadRequest)

		return
	}

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

	key := user.Key(register.Email)
	exists, err := storage.GetInstance().Exist().Entity(key).NowT()

	if err != nil {
		logger.Infof("Could not check if key %s exists due to: %s", register.Email, err.Error())
	}

	if exists {
		msg := "User with this email already exists."

		logger.Info(msg)

		http.Error(writer, msg, http.StatusBadRequest)

		return
	}

	user := user.New(register.Name, register.Email, register.Password)
	save_err := storage.GetInstance().Save().Entity(user).Now()

	if save_err != nil {
		http.Error(writer, "Could not register user.", http.StatusInternalServerError)

		return
	}

	logger.Info("Successful register")
}
