package main

import (
	"net/http"
	"time"
)

func deletecookie(w http.ResponseWriter, r *http.Request) {
	cook, _ := r.Cookie("user")
	c := http.Cookie{
		Name:    cook.Value,
		MaxAge:  -1,
		Expires: time.Now(),
	}
	http.SetCookie(w, &c)

	/*
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "user", Value: "null", Expires: expiration}
		http.SetCookie(w, &cookie)
	*/
	http.Redirect(w, r, "/", http.StatusFound)
}
