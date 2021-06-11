package main

import (
	"fmt"
	"net/http"
)

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, cookie.Value)
	}

}
