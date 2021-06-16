package main

import (
	"html/template"
	"net/http"
)

func handle_login(w http.ResponseWriter, r *http.Request) {
	//	v := 1
	username := r.FormValue("email")
	pass := r.FormValue("pass")
	tmp, _ := template.ParseFiles("public/login/index.html")
	tmp.Execute(w, nil)

	date := login(username, pass)
	if date == "null" {
		return
	}

}

func login(username, pass string) string {
	if username == "" || pass == "" {
		return "null"
	}

	user := select_user(username, pass)
	return user
}
