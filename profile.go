package main

import (
	"html/template"
	"net/http"
)

func handle_perfil(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("public/profile/index.html")
	tmp.Execute(w, nil)

	/*cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}


	fmt.Println("Interfaz de perfil de usuario 1.01")
	fmt.Fprintf(w, "<h1>%s</h1>", cookie.Value)
	*/

}
