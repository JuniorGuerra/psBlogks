package main

import (
	"html/template"
	"net/http"
)

func recuperate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/recuperate_pass/pass.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}
