package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/user"
	"github.com/ilker-raimov/cca/primary/jwt"
	"github.com/ilker-raimov/cca/primary/util/regex"

	logger "github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
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

	token, err := jwt.Create(user.Email)

	if err != nil {
		http.Error(writer, "Could not create JWT.", http.StatusInternalServerError)

		return
	}

	response := LoginResponse{
		Token: token,
	}
	response_data, err := json.Marshal(response)

	if err != nil {
		http.Error(writer, "Could not map JWT.", http.StatusInternalServerError)

		return
	}

	writer.Write(response_data)
}

func Register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var register RegisterRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&register); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	logger.Infof("Username: %s", register.Username)
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

	is_email_valid := regex.Email(register.Email)

	if !is_email_valid {
		http.Error(writer, "Invalid email.", http.StatusBadRequest)

		return
	}

	is_password_valid := regex.Password(register.Password)

	if !is_password_valid {
		http.Error(writer, "Invalid password. It must contain 1 capital letter, 1 digit, 1 special symbol and be atleast 8 symbols.", http.StatusBadRequest)

		return
	}

	user := user.New(register.Username, register.Email, register.Password)
	save_err := storage.GetInstance().Save().Entity(user).Now()

	if save_err != nil {
		http.Error(writer, "Could not register user.", http.StatusInternalServerError)

		return
	}

	logger.Info("Successful register")
}
