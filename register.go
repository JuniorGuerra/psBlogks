package main

import (
	"html/template"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {

	f := r.FormValue("d")
	if f == "user" {
		w.Write([]byte("<script>alert('Los nombre de usuario existente')</script>"))
	}

	if f == "mail" {
		w.Write([]byte("<script>alert('Mail existente')</script>"))
	}

	if f == "e_mail" {
		w.Write([]byte("<script>alert('Error al verificar el mail')</script>"))
	}

	if f == "fail" {
		w.Write([]byte("<script>alert('Datos de usuario existentes intente cambiar nombre de usuario o correo')</script>"))
	}

	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
}
