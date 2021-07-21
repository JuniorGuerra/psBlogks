package main

import (
	"html/template"
	"net/http"
)

func handle_edit_profile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("public/edit_profile/index.html")
	tmpl.Execute(w, nil)
}

func verify_edit_profile(w http.ResponseWriter, r *http.Request) {
	img := r.FormValue("img")
	resume := r.FormValue("resume")

	/*	if img != "" && resume != "" {
		//Errores verificar aqui
	}*/

	//Esta funcion es para almacenar la imagen en nuestra database

	img2html := "<html><body><img src=\"data:image/png;base64," + img + "\" /></body></html>"

	w.Write([]byte(img2html))
	w.Write([]byte(resume))
}

func handle_edit_book(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>In this page is for edit book</h1>"))
}
