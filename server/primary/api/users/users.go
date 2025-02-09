package users

import (
	"net/http"

	"github.com/ilker-raimov/cca/common/storage/model/model_user"
	"github.com/ilker-raimov/cca/common/util/response"
)

func Roles(writer http.ResponseWriter, request *http.Request) {
	roles := model_user.Roles()

	response.WriteOrInternal(writer, roles)
}
