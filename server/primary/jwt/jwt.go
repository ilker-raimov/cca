package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilker-raimov/cca/common/environment"
	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	logger "github.com/sirupsen/logrus"
)

var ErrInvalidToken = errors.New("invalid token")

func Create(email string, role model_user.Role) (string, error) {
	expire_at := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   expire_at,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sign_key := environment.GetOrPanic("JWT_SIGN_KEY")

	return token.SignedString([]byte(sign_key))
}

func Parse(data string) (map[string]interface{}, error) {
	logger.Infof("Parsing: %s", data)

	sign_key := environment.GetOrPanic("JWT_SIGN_KEY")
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(sign_key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func ParseAndVerify(data string, roles []model_user.Role, writer http.ResponseWriter) (string, bool) {
	claims, err := Parse(data)

	if err != nil {
		logger.Errorf("Could not check JWT due to: %s", err)

		http.Error(writer, "Could not check JWT.", http.StatusInternalServerError)

		return "", false
	}

	_, is_exp_ok := parseAndVerifyExp(claims, writer)

	if !is_exp_ok {
		return "", false
	}

	_, is_role_ok := parseAndVerifyRole(claims, roles, writer)

	if !is_role_ok {
		return "", false
	}

	return parseAndVerifyEmail(claims, writer)
}

func parseAndVerifyExp(claims map[string]interface{}, writer http.ResponseWriter) (int64, bool) {
	return parseAndVerifyT("exp", -1, func(to_map interface{}) (int64, bool) {
		mapped, ok := to_map.(float64)

		return int64(mapped), ok
	}, func(value int64) bool {
		return value > time.Now().Unix()
	}, func() {
		http.Error(writer, "Expired JWT.", http.StatusUnauthorized)
	}, claims, writer)
}

func parseAndVerifyRole(claims map[string]interface{}, desired_roles []model_user.Role, writer http.ResponseWriter) (int, bool) {
	return parseAndVerifyT("role", -1, func(to_map interface{}) (int, bool) {
		mapped, ok := to_map.(float64)

		return int(mapped), ok
	}, func(value int) bool {
		for _, role := range desired_roles {
			if int(role) == value {
				return true
			}
		}

		return false
	}, func() {
		http.Error(writer, "Not enough permissions.", http.StatusForbidden)
	}, claims, writer)
}

func parseAndVerifyEmail(claims map[string]interface{}, writer http.ResponseWriter) (string, bool) {
	return parseAndVerifyT("email", "", func(to_map interface{}) (string, bool) {
		mapped, ok := to_map.(string)

		return mapped, ok
	}, func(value string) bool {
		return true
	}, func() {}, claims, writer)
}

func parseAndVerifyT[T any](name string, fallback T, mapper func(interface{}) (T, bool), checker func(T) bool, failed func(), claims map[string]interface{}, writer http.ResponseWriter) (T, bool) {
	logger.Infof("Parsing and verifying: %s", name)

	data, exists := claims[name]

	if !exists {
		message := fmt.Sprintf("Invalid JWT. Missing %s property.", name)

		logger.Warn(message)

		http.Error(writer, message, http.StatusBadRequest)

		return fallback, false
	}

	value, ok := mapper(data)

	if !ok {
		data_type := reflect.TypeOf(fallback)
		type_name := data_type.Name()
		message := fmt.Sprintf("Invalid JWT. Property %s is not of type %s.", name, type_name)

		logger.Warn(message)

		http.Error(writer, message, http.StatusBadRequest)

		return fallback, false
	}

	is_ok := checker(value)

	if !is_ok {
		logger.Warn("Property is not valid.")

		failed()

		return fallback, false
	}

	logger.Infof("Property %s with value %v is ok.", name, value)

	return value, true
}
