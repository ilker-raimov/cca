package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/log"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, "Login requires POST", http.StatusMethodNotAllowed)

		return
	}

	defer request.Body.Close()

	var login LoginRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&login); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	log.InfoF("Email: %s", login.Email)
	log.InfoF("Password: %s", login.Password)

	if login.Email == "fail" || login.Password == "fail" {
		log.Info("Unsuccessful login")

		http.Error(writer, "Invalid credentials", http.StatusBadRequest)

		return
	}

	log.Info("Successful login")
}
