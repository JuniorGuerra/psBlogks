package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type viewDataBook struct {
	Autor    string
	Title    string
	Created  string
	Category string
}

func handle_view_blog(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("public/read_blog/index.html")
	if err != nil {
		panic("Error: " + err.Error())
	}

	vars := mux.Vars(r)
	user := vars["user"]
	title := vars["title"]
	//El blog y el book hacen referencia  a lo mismo literalmente
	viewData := selectBookUser(user, title)
	Data := viewDataBook{
		Autor:    viewData.autor,
		Title:    viewData.title,
		Created:  viewData.fecha,
		Category: viewData.categoria,
	}
	tmp.Execute(w, Data)

	scriptHtml := "<div style='margin-right: 250px; margin-left: 100px'>" + viewData.body + "</div>"
	w.Write([]byte(scriptHtml))
}
