package main

import (
	"html/template"
	"net/http"
	"time"
)

func handle_login(w http.ResponseWriter, r *http.Request) {
	//	v := 1
	username := r.FormValue("email")
	pass := r.FormValue("pass")
	tmp, _ := template.ParseFiles("public/login/index.html")
	tmp.Execute(w, nil)

	date := login(username, pass)
	if date == "null" {
		w.Write([]byte("<script>alert('Ingrese todos los datos.')</script>"))
		return
	} else {

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/profile", http.StatusFound)

	}
}

func login(username, pass string) string {
	if username == "" || pass == "" {
		return "null"
	}

	user := select_user(username, pass)
	return user
}
