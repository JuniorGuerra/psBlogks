package main

import (
	"html/template"
	"net/http"
)

type data_user struct {
	Name     string
	Username string
	Resumen  string
	Email    string
}

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("public/profile/index.html")

	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	name_user, resume_user := select_user_view(cookie.Value)

	user_data := data_user{
		Name:     "Perfil: " + cookie.Value,
		Username: name_user,
		Resumen:  resume_user,
		Email:    "email@example.com",
	}
	tmp.Execute(w, user_data)

}
