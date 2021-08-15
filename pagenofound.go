package main

import (
	"html/template"
	"net/http"
)

// Handle404 ...
func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("public/Error404/index.html")

		if err != nil {
			panic(err.Error)
		}

		t.Execute(w, nil)
	})
}
