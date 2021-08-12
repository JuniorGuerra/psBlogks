package main

import (
	"fmt"
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
	query_users := select_users_query_all(name)
	for i, val := range query_users {
		fmt.Println(i, val)
	}

	fmt.Println("vista desde query", query_users[1])
	t.Execute(w, d)
}
