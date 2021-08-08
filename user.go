package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	img, name, phone, description := select_user_view(user)

	if name == "" && description == "" {
		w.Write([]byte("<h1>Usuario con informacion no publica</h1>"))
		return

	}
	w.Write([]byte("<h1>Image: </h1>" + img))
	w.Write([]byte("<h1>Phone: </h1>" + phone))
	w.Write([]byte("<h1>usuario: </h1>" + name))
	w.Write([]byte("<h2>descripcion: </h2>" + description))

}
