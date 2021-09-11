package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type datos_blog struct {
	Title     string
	Autor     string
	Creacion  string
	Categoria string
	Url       string
}

func handler_buscar(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("query") != "" {
		http.Redirect(w, r, "/blogs/"+r.FormValue("query"), http.StatusFound)
		return
	}

	t, err := template.ParseFiles("public/buscar_blog/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, nil)
}

func handler_blogs(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("public/buscar_blog/lista_blog.html")
	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)
	blog := vars["blogs"]

	if blog == "" {
		http.Redirect(w, r, "/explorar", http.StatusFound)
		return
	}

	d_blog := selectBooks(blog)

	if d_blog == nil {
		w.Write([]byte("No se ha encontrado un blog con ese nombre"))
		return
	}
	var exist int
	for _, v := range d_blog {
		datos := datos_blog{
			Title:     v.title,
			Autor:     v.autor,
			Creacion:  v.fecha,
			Categoria: v.categoria,
			Url:       v.autor + "/blog/" + v.title,
		}
		t.Execute(w, datos)
		exist++
	}

	if exist == 0 {
		w.Write([]byte("No se ha encontrado un blog con ese nombre"))
	}

}
