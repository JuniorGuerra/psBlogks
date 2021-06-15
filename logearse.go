package main

import (
	"html/template"
	"net/http"
)

func handle_login(w http.ResponseWriter, r *http.Request) {
	//	var username string
	//	var pass string

	//	username = "juniorguerrac17@gmail.com"
	//	pass = "Prueba2021"

	tmp, _ := template.ParseFiles("public/login/index.html")
	tmp.Execute(w, nil)
}

func login(username, pass string) (string, string) {
	if username == "" || pass == "" {
		return "Ingrese todos los datos por favor", ""
	}

	user, mail := select_user(username, pass)
	return user, mail
}
