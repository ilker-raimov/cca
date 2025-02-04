package run

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ilker-raimov/cca/common/util/response"
	"github.com/ilker-raimov/cca/docker/code"
)

type CompileResponse struct {
	Ok     bool
	Output string
}

func Compile(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	data, err := io.ReadAll(request.Body)

	if err != nil {
		response.InternalServerError(writer, "Failed to read request")

		return
	}

	java := code.New()
	ok, output, err := java.Compile(data)

	if err != nil {
		response.InternalServerError(writer, output)

		return
	}

	compile_response := CompileResponse{Ok: ok, Output: output}

	if data, err := json.Marshal(compile_response); err != nil {
		response.InternalServerError(writer, "Failed to serialize response")
	} else {
		writer.Write(data)
	}
}

func Check(writer http.ResponseWriter, request *http.Request) {

}

func Submit(writer http.ResponseWriter, request *http.Request) {

}
