package main

import (
	"fmt"
	"net/http"
)

func deletecookie(w http.ResponseWriter, r *http.Request) {
	cook, _ := r.Cookie("user")
	fmt.Fprintf(w, "%s", cook.Value)
}
