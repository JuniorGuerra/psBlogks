package main

import (
	"net/http"
)

// Handle404 ...
func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola esta pagina esta caida"))
	})
}