package main

import (
	"html/template"
	"net/http"
)

// Handle404 ...
func pstatic(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")

		if err != nil {
			panic(err)
		}

		t.Execute(w, r)
}
