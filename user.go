package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type dates struct {
	Name        string
	Description string
	Phone       string
}

func user(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("public/users/index.html")

	if err != nil {
		panic(err.Error)
	}

	vars := mux.Vars(r)
	user := vars["user"]

	img, name, phone, description := select_user_view(user)
	d := dates{
		Name:        name,
		Description: description,
		Phone:       phone,
	}
	if name == "" && description == "" {
		w.Write([]byte("<h1>Usuario con informacion no publica</h1>"))
		return

	}
	var text string = "<div><img src=\"data:image/png;base64," + img + "\" alt='Imagen profile user' class='perfil_img' id='img'></div>"
	w.Write([]byte(text))
	t.Execute(w, d)

}
