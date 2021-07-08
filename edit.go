package main

import "net/http"

func handle_edit_profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>In this page is for edit profile</h1>"))
}

func handle_edit_book(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>In this page is for edit book</h1>"))
}
