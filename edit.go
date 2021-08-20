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
	phone := r.FormValue("phone")

	if img == "" || resume == "" {
		http.Redirect(w, r, "/edit/profile", http.StatusFound)
		return
	}

	//img2html := "<html><body><img class="img_profile" alt="imagen_perfil_usuario" src=\"data:image/png;base64," + img + "\" /></body></html>"
	name, err := r.Cookie("user")

	if err != nil {
		panic(err)
	}

	update_data_profile(img, name.Value, phone, resume)

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func handle_edit_book(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>In this page is for edit book</h1>"))
}
