package main

import (
	"html/template"
	"net/http"
)

type data_user struct {
	Name     string
	Username string
	Phone    string
	Resumen  string
	Email    string
}

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("public/profile/index.html")

	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	image_user, name_user, phone_user, resume_user := select_user_view(cookie.Value)
	for i := 0; i < 30; i++ {
		//Este es un for vacio, pero es para que no me ejecute los 2 datos al mismo tiempo y asi evitar un posible error de tcp
	}
	gmail_user := select_gmail(cookie.Value)

	if resume_user == "" {
		resume_user = "Aun no tienes informacion para mostrar"
	}
	//image := "\"data:image/png;base64," + image_user + "\" "
	user_data := data_user{
		Name:     "Perfil: " + cookie.Value,
		Username: name_user,
		Phone:    phone_user,
		Resumen:  resume_user,
		Email:    gmail_user,
	}
	var text string = "<div><img src=\"data:image/png;base64," + image_user + "\" alt='Imagen profile user' class='perfil_img'></div>"

	w.Write([]byte(text))
	tmp.Execute(w, user_data)

}
