package main

import (
	"fmt"
	"net/http"
)

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Println("Interfaz de perfil de usuario 1.01")
	fmt.Fprintf(w, "<h1>%s</h1>", cookie.Value)

}
