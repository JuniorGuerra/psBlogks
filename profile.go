package main

import (
	"html/template"
	"net/http"
)

type data_user struct {
	Name      string
	Username  string
	Phone     string
	Resumen   string
	Email     string
	Books     int
	NameBooks string
}

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("public/profile/index.html")

	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if cookie.Value == "null" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	image_user, _, phone_user, resume_user := select_user_view(cookie.Value)

	gmail_user := select_gmail(cookie.Value)
	books_user := selectAllBookUser(cookie.Value)
	var x int
	var nameBook []string
	for _, v := range books_user {
		nameBook = append(nameBook, v.title)
		x++
	}

	if resume_user == "" {
		resume_user = "Aun no tienes informacion para mostrar, si creas un blog, pero no tienes tu perfil editado, no se podra ver tu perfil"
	}
	//image := "\"data:image/png;base64," + image_user + "\" "
	user_data := data_user{
		Name:      "Perfil: " + cookie.Value,
		Username:  cookie.Value,
		Phone:     phone_user,
		Resumen:   resume_user,
		Email:     gmail_user,
		Books:     x,
		NameBooks: "Presione uno para editar",
	}

	var text string = "<a href='/user/" + cookie.Value + "'><img class='perfil_img' src=\"data:image/png;base64," + image_user + "\" alt='Imagen profile user'>"

	tmp.Execute(w, user_data)
	w.Write([]byte("<div style='margin-top:-30px'>"))
	for _, v := range nameBook {
		w.Write([]byte("<p style='margin-left:7%; font-size:1.5rem;' >" + v + "</p>"))
	}
	w.Write([]byte("</div>"))

	w.Write([]byte(text))
}
