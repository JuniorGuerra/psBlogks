package main

import (
	"fmt"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {
	var name string
	var email string
	var pass string

	result := register(name, email, pass)
	fmt.Println(result)
}

func register(username, email, pass string) string {
	if pass == "" || username == "" || email == "" {
		return "usuario malo"
	}
	return insert_new_user(username, email, pass)
}
