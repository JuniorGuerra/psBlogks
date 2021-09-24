package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func handle_edit_profile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("public/edit_profile/index.html")
	tmpl.Execute(w, nil)
}

func verify_edit_profile(w http.ResponseWriter, r *http.Request) {
	img := r.FormValue("img")
	resume := r.FormValue("resume")
	phone := r.FormValue("phone")

	if img == "" || resume == "" {
		http.Redirect(w, r, "/edit/profile", http.StatusFound)
		return
	}

	//img2html := "<html><body><img class="img_profile" alt="imagen_perfil_usuario" src=\"data:image/png;base64," + img + "\" /></body></html>"
	name, err := r.Cookie("user")

	if err != nil {
		panic(err)
	}

	update_data_profile(img, name.Value, phone, resume)

	http.Redirect(w, r, "/profile", http.StatusFound)
}

var id string

func handle_edit_book(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]
	title := vars["title"]
	//El blog y el book hacen referencia  a lo mismo literalmente
	viewData := selectBookUser(user, title)
	id = viewData.id
	w.Write([]byte(fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<script src="https://kit.fontawesome.com/a32a015951.js"></script>
		<link rel="shortcut icon" href="Img/logo_blogbook.png">
	
		<title>BlogBook-Crear</title>
	</head>
	
	<body>
		<script src="//cdn.ckeditor.com/4.16.2/full/ckeditor.js"></script>
		<form method="post" action="/update">
			<textarea name="editor1" id="editor1" rows="10" cols="80">      
				%s
			</textarea>
			<script>
				CKEDITOR.replace('editor1');
			</script>
			<input type="submit" value="Publicar">
		</form>
		</script>
	</body>
	
	</html>
	`, viewData.body)))
}

func handler_update(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("editor1")

	if len(text) < 10 {
		http.Redirect(w, r, "/create", http.StatusBadRequest)
	}
	//El blog y el book hacen referencia  a lo mismo literalmente
	if !update_book(id, text) {
		w.Write([]byte("<script>alert('No se logro actualizar el blog')</script>"))
		return
	}
	http.Redirect(w, r, "/explorar", http.StatusFound)
}
