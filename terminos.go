package main

import (
	"html/template"
	"net/http"
)

func terminos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/terminos/index.html")
	if err != nil {
		panic(err.Error)
	}

	t.Execute(w, nil)
}
