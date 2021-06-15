package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func handle_register(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("password")
	repass := r.FormValue("rpassword")
	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
	fmt.Println(name + email + pass + repass)

	if pass != repass {
		w.Write([]byte("<script>alert('las contraseñas no coinciden')</script>"))
		return
	}

	result := register(name, email, pass)

	w.Write([]byte("<script>alert('" + result + "')</script>"))

	if result == "error" {
		w.Write([]byte("<script>alert('Los datos de usuario son existente')</script>"))
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: name, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/profile", http.StatusFound)

}

func register(username, email, pass string) string {
	if pass == "" || username == "" || email == "" {
		return "usuario malo"
	}
	return insert_new_user(username, email, pass)
}
