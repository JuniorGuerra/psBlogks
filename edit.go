package main

import (
	"html/template"
	"net/http"
)

func handle_edit_profile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("public/edit_profile/index.html")
	tmpl.Execute(w, nil)
}

func handle_edit_book(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>In this page is for edit book</h1>"))
}
