package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, "Login requires POST", http.StatusMethodNotAllowed)

		return
	}

	defer request.Body.Close()

	fmt.Println("Login:")

	var login LoginRequest

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&login); err != nil {
		http.Error(writer, "Invalid request format.", http.StatusBadRequest)

		return
	}

	fmt.Printf("Email: %s\n", login.Email)
	fmt.Printf("Password: %s\n", login.Password)

	if login.Email == "fail" || login.Password == "fail" {
		fmt.Println("Unsuccessful login")

		http.Error(writer, "Invalid credentials", http.StatusBadRequest)

		return
	}

	fmt.Println("Successful login")
}
