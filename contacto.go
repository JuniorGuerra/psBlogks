package main

import (
	"html/template"
	"net/http"
)

func contactenos(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("e")

	if f == "excelente" {
		w.Write([]byte("<script>alert('Mensaje enviado correctamente')</script>"))
	}

	if f == "error" {
		w.Write([]byte("<script>alert('Ha habido un erro al contactarnos')</script>"))
	}

	if f == "vacio" {
		w.Write([]byte("<script>alert('Llene todos los campos para poder reconocerlo')</script>"))
	}

	tmp, err := template.ParseFiles("public/contactos/index.html")
	if err != nil {
		panic(err)
	}
	tmp.Execute(w, nil)
}
