package interceptor

import (
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func LogInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("Request: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
