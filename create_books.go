package main

import (
	"fmt"
	"net/http"
)

func handle_create(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")

	if cookie == nil {
		w.Write([]byte("<h1>Creese un usuario para escribir un libro</h1>"))
		return
	}
	fmt.Fprintf(w, cookie.Value)

	title := "El cantero del diablo"
	body := "El principe murio en 2021 y el cantero lo violo"
	author := cookie.Value
	category := "fantasia"

	result := insert_new_book(title, body, author, category)
	fmt.Println(result)

}
