package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type dates struct {
	Name        string
	Description string
	Phone       string
	Books       int32
}

func user(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("public/users/index.html")

	if err != nil {
		panic(err.Error)
	}

	vars := mux.Vars(r)
	user := vars["user"]

	img, name, phone, description := select_user_view(user)

	books_user := selectAllBookUser(name)
	var x int32

	for range books_user {
		x++
	}

	d := dates{
		Name:        name,
		Description: description,
		Phone:       phone,
		Books:       x,
	}

	if name == "" && description == "" {
		w.Write([]byte("<h1>Usuario con informacion no publica</h1>"))
		return
	}
	var text string = "<div><img src=\"data:image/png;base64," + img + "\" alt='Imagen profile user' class='perfil_img' id='img'></div>"
	w.Write([]byte(text))
	t.Execute(w, d)
	w.Write([]byte("<div style='margin-left:70px; color: black;'>"))
	for _, v := range books_user {
		a := fmt.Sprintf("<a href='/%s/blog/%s'><h3>%s</h3></a> <p> Creado el: %s</p>", name, v.title, v.title, v.fecha)
		w.Write([]byte(a))
	}
	w.Write([]byte("</div>"))
}
