package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	fmt.Fprintf(w, user)

	if user != "Junior" {
		fmt.Fprintf(w, "Usuario no existe")
	}

}
