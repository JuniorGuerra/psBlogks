package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	var dir string
	flag.StringVar(&dir, "public/", ".", "/public/")
	flag.Parse()

	r := mux.NewRouter()
	r.NotFoundHandler = Handle404()

	//Los archivos estaticos estan en localhost:port/static/nombre
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/", home)
	r.HandleFunc("/register", handle_register)
	r.HandleFunc("/login", handle_login)
	r.HandleFunc("/profile", handle_perfil)
	r.HandleFunc("/static/", pstatic)
	r.HandleFunc("/user/{user}", user)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
		//Buenas practicas de espera para los servicios
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server -> localhost:" + port)
	log.Fatal(srv.ListenAndServe())

}

//Pagina del home
func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}
	tmp.Execute(w, "Helpme")
}
