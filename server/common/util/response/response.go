package response

import (
	"encoding/json"
	"io"
	"net/http"
)

func BadRequest(writer http.ResponseWriter, error string) {
	http.Error(writer, error, http.StatusBadRequest)
}

func InternalServerError(writer http.ResponseWriter, error string) {
	http.Error(writer, error, http.StatusInternalServerError)
}

func ParseOrInternal(writer http.ResponseWriter, reader io.ReadCloser, object any) bool {
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(&object); err != nil {
		BadRequest(writer, "Invalid request format.")

		return false
	}

	return true
}

func WriteOrInternal(writer http.ResponseWriter, data any) {
	marshal_data, err := json.Marshal(data)

	if err != nil {
		InternalServerError(writer, "Could not marshal response.")

		return
	}

	writer.Write(marshal_data)
}
