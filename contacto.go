package main

import (
	"html/template"
	"net/http"
)

func contactenos(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("public/contactos/index.html")
	if err != nil {
		panic(err)
	}
	tmp.Execute(w, nil)
}
