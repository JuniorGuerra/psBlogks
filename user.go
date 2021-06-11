package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	name, des := select_user_view(user)
	w.Write([]byte("<h1>usuario: </h1>" + name))
	w.Write([]byte("<h2>descripcion: </h2>" + des))

}
