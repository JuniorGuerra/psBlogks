package main

import (
	"html/template"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
}
