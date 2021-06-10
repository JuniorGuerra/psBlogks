package main

import (
	"fmt"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {
	value := register("", "juniorguerrac17@gmail.com", "JuniorEnPrueba")
	fmt.Println(value)
}

func register(username, email, pass string) string {
	if pass == "" || username == "" || email == "" {
		return "usuario malo"
	}

	return "usuario disponible para el registro"

}
