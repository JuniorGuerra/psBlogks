package main

import (
	"html/template"
	"net/http"
)

func handle_create(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")

	if cookie == nil {
		w.Write([]byte("<h1>Creese un usuario para escribir un libro</h1>"))
		return
	}
	t, err := template.ParseFiles("public/create/index.html")

	if err != nil {
		panic(err.Error())
	}

	t.Execute(w, nil)
	/*title := "El cantero del diablo"
	body := "El principe murio en 2021 y el cantero lo violo"
	author := cookie.Value
	category := "fantasia"

	result := insert_new_book(title, body, author, category)
	fmt.Println(result)*/

}

func HandlePublicar(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("editor1")

	if len(text) < 10 {
		http.Redirect(w, r, "/create", http.StatusBadRequest)
	}

	cookie, _ := r.Cookie("user")

	if cookie == nil {
		w.Write([]byte("<h1>Ya probo el sistema, ahora crea un usuario para escribir un libro</h1>"))
		return
	}

	w.Write([]byte(insert_new_book("Titulo", text, cookie.Value, "Prueba")))

}
