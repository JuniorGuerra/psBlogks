package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	name, description := select_user_view(user)

	if name == "" && description == "" {
		w.Write([]byte("<h1>Usuario con informacion no publica</h1>"))
		return
	}
	w.Write([]byte("<h1>usuario: </h1>" + name))
	w.Write([]byte("<h2>descripcion: </h2>" + description))

}
