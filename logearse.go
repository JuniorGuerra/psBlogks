package main

import (
	"html/template"
	"net/http"
)

func handle_login(w http.ResponseWriter, r *http.Request) {
	//	v := 1
	f := r.FormValue("d")

	if f == "fail" {
		w.Write([]byte("<script>alert('Usuario o contraseña incorrectos, intente nuevamente ')</script>"))
	}

	if f == "exitoso" {
		w.Write([]byte("<script>alert('Contraseña cambiada exitosamente ')</script>"))
	}

	tmp, _ := template.ParseFiles("public/login/index.html")
	tmp.Execute(w, nil)
}
