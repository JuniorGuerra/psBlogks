package main

import (
	"net/http"
)

func sendMessage(w http.ResponseWriter, r *http.Request) {

	fullName := r.FormValue("fullname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	asunto := r.FormValue("affair")
	message := r.FormValue("message")
	if fullName != "" && email != "" && message != "" {

		if !Sendemail(fullName, email, phone, asunto, message) {
			http.Redirect(w, r, "/Contactenos?e=error", http.StatusFound)
			return
		}

		http.Redirect(w, r, "/Contactenos?e=excelente", http.StatusFound)
		return
	} else {

		http.Redirect(w, r, "/Contactenos?e=vacio", http.StatusFound)
		return

	}

}
