package auth

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/storage"
	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	"github.com/ilker-raimov/cca/common/util/response"
	"github.com/ilker-raimov/cca/primary/jwt"
	"github.com/ilker-raimov/cca/primary/util/regex"

	logger "github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Role     model_user.Role `json:"role"`
	Token    string          `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var login LoginRequest

	is_parse_ok := response.ParseOrInternal(writer, request.Body, &login)

	if !is_parse_ok {
		return
	}

	logger.Infof("Email: %s", login.Email)
	logger.Infof("Password: %s", login.Password)

	key := model_user.Key(login.Email)
	exists, err := storage.GetInstance().Exist().Entity(key).NowT()

	if err != nil {
		response.InternalServerError(writer, model_user.COULD_NOT_CHECK)

		return
	}

	if !exists {
		response.BadRequest(writer, model_user.NO_SUCH_USER)

		return
	}

	var user model_user.User

	load_err := storage.GetInstance().Load().Entity(&user, key).Now()

	if load_err != nil {
		response.InternalServerError(writer, model_user.COULD_NOT_LOAD)

		return
	}

	match := user.Password == login.Password

	if !match {
		response.BadRequest(writer, "Invalid password.")

		return
	}

	token, err := jwt.Create(user.Email, user.Role)

	if err != nil {
		response.InternalServerError(writer, "Could not create JWT.")

		return
	}

	login_response := LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	response.WriteOrInternal(writer, login_response)
	logger.Info("Successful login")
}

func Register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var register RegisterRequest

	is_parse_ok := response.ParseOrInternal(writer, request.Body, &register)

	if !is_parse_ok {
		return
	}

	logger.Infof("Username: %s", register.Username)
	logger.Infof("Email: %s", register.Email)
	logger.Infof("Password: %s", register.Password)

	key := model_user.Key(register.Email)
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

	user := model_user.New(register.Username, register.Email, register.Password)
	save_err := storage.GetInstance().Save().Entity(user).Now()

	if save_err != nil {
		http.Error(writer, "Could not register user.", http.StatusInternalServerError)

		return
	}

	logger.Info("Successful register")
}
