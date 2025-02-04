package response

import "net/http"

func BadRequest(writer http.ResponseWriter, error string) {
	http.Error(writer, error, http.StatusBadRequest)
}

func InternalServerError(writer http.ResponseWriter, error string) {
	http.Error(writer, error, http.StatusInternalServerError)
}
