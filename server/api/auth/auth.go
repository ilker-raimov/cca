package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/log"
	"github.com/ilker-raimov/cca/storage"
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

	log.InfoF("Email: %s", login.Email)
	log.InfoF("Password: %s", login.Password)

	storage.GetInstance().Load(login.Email)

	log.Info("Successful login")
}

func Register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var register RegisterRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&register); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	log.InfoF("Name: %s", register.Name)
	log.InfoF("Email: %s", register.Email)
	log.InfoF("Password: %s", register.Password)

	exists, err := storage.GetInstance().Exists(register.Email)

	if err != nil {
		log.InfoF("Could not check if key %s exists due to: %s", register.Email, err.Error())
	}

	if exists {
		msg := "User with this email already exists."

		log.Info(msg)

		http.Error(writer, msg, http.StatusBadRequest)

		return
	}

	storage.GetInstance().Save(register.Email, []byte(register.Name))

	log.Info("Successful register")
}
