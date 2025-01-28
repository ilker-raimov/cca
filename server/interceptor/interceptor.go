package interceptor

import (
	"net/http"

	"github.com/ilker-raimov/cca/log"
)

func LogInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.InfoF("Request: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
