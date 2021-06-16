package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("password")
	repass := r.FormValue("rpassword")
	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
	fmt.Println(name + email + pass + repass)
}
