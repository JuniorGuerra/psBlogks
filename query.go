package main

import (
	"html/template"
	"net/http"
)

type dato struct {
	Query string
}

func query(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("buscar")
	var d dato = dato{Query: name}

	t, err := template.ParseFiles("public/query/index.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, d)

	query_users := select_users_query_all(name)
	w.Write([]byte("<div class='text' id='text'>"))
	for _, val := range query_users {
		w.Write([]byte("<a href='https://bsblogbook.herokuapp.com/user/" + val.name + "'><h2>" + val.name + "</h2></a>"))
	}
	w.Write([]byte("</div>"))
}
