package main

import (
	"html/template"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {

	f := r.FormValue("d")
	if f == "fail" {
		w.Write([]byte("<script>alert('Los datos de usuario ya existen')</script>"))
	}

	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
}
