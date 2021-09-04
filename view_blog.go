package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type viewData struct {
	autor     string
	title     string
	tipo      string
	fecha     string
	categoria string
	id        string
}

func handle_view_blog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	title := vars["title"]

	//El blog y el book hacen referencia  a lo mismo literalmente
	viewData := selectBookUser(user, title)

	w.Write([]byte(viewData.body))

}
