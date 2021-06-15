package main

import (
	"html/template"
	"net/http"
)

func handle_register(w http.ResponseWriter, r *http.Request) {
	/*
		var name string
		var email string
		var pass string
		name = "Kevin"
		email = "Kevin@blogbook.com"
		pass = "Kevin2021"

		result := register(name, email, pass)
		fmt.Println(result)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: name, Expires: expiration}
		http.SetCookie(w, &cookie)
	*/
	tmp, _ := template.ParseFiles("public/register/index.html")
	tmp.Execute(w, nil)
}

func register(username, email, pass string) string {
	if pass == "" || username == "" || email == "" {
		return "usuario malo"
	}
	return insert_new_user(username, email, pass)
}
