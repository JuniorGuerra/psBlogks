package main

import (
	"net/http"
)

func deletecookie(w http.ResponseWriter, r *http.Request) {
	cook, _ := r.Cookie("user")
	c := http.Cookie{
		Name:   cook.Value,
		MaxAge: -1}
	http.SetCookie(w, &c)

	http.Redirect(w, r, "/", http.StatusFound)
}
