package interceptor

import (
	"fmt"
	"net/http"
)

func LogInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
