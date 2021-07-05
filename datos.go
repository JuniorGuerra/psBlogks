package main

import (
	"fmt"
	"net/http"
	"time"
)

func datos_login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("email")
	pass := r.FormValue("pass")

	if name == "" || pass == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	name = select_user(name, pass)

	if name == "noUser" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Println(name)
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "user", Value: name, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)

}

func datos_register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("password")
	fmt.Println(name, email, pass)
	value := insert_new_user(name, email, pass)

	if !value {
		http.Redirect(w, r, "/register", http.StatusNotFound)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "user", Value: name, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)
}
