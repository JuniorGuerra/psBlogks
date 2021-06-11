package main

import (
	"fmt"
	"net/http"
)

func handle_login(w http.ResponseWriter, r *http.Request) {
	var username string
	var pass string

	username = "juniorguerrac17@gmail.com"
	pass = "Prueba2021"

	user, mail := login(username, pass)
	if user == "" || mail == "" {
		fmt.Println("Usuario no registrado")
	}
	fmt.Println(user + mail)
	http.Redirect(w, r, "user/"+user, http.StatusFound)

}

func login(username, pass string) (string, string) {
	if username == "" || pass == "" {
		return "Ingrese todos los datos por favor", ""
	}

	user, mail := select_user(username, pass)
	return user, mail
}
