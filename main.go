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
	r.HandleFunc("/edit/profile", handle_edit_profile)
	r.HandleFunc("/edit/book", handle_edit_book)
	r.HandleFunc("/verify_edit_profile", verify_edit_profile)
	r.HandleFunc("/static/", pstatic)
	r.HandleFunc("/verify", datos_login)
	r.HandleFunc("/verifynew", datos_register)
	r.HandleFunc("/verify_mail", handlerMail)
	r.HandleFunc("/register/verify_mail", dMail)
	r.HandleFunc("/user", query)
	r.HandleFunc("/user/{user}", user)
	r.HandleFunc("/delcook", deletecookie)
	r.HandleFunc("/about", terminos)
	r.HandleFunc("/Contactenos", contactenos)

	r.HandleFunc("/exit", deletecookie)

	//Recuperate password
	r.HandleFunc("/pass", recuperate)
	r.HandleFunc("/recuperate_pass", handleRecuperate)
	r.HandleFunc("/verify_code", handleVerifyCode)
	r.HandleFunc("/Dverify_code", verify_code)
	r.HandleFunc("/change_password", handleChangePassword)
	r.HandleFunc("/verify_change_password", change_password)

	//Enviarnos mensaje
	r.HandleFunc("/sendMessage", sendMessage)

	r.HandleFunc("/create", handle_datos_create)
	r.HandleFunc("/bnew/create", handle_create)
	r.HandleFunc("/publicar", HandlePublicar)
	r.HandleFunc("/{user}/blog/{title}", handle_view_blog)

	//Buscar libros
	r.HandleFunc("/explorar", handler_buscar)
	r.HandleFunc("/blogs/{blogs}", handler_blogs)
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
		//Buenas practicas de espera para los servicios
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("http://localhost:" + port + "/")
	log.Fatal(srv.ListenAndServe())

}

//Pagina del home

type d struct {
	Ir string
}

var a string

func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Println("No se encuentra el archivo")
	}

	cookie, err := r.Cookie("user")
	if err != nil {
		a = "Iniciar"
	} else {
		if cookie.Value == "null" {
			a = "Iniciar"
		} else {
			a = "Perfil"
		}
	}

	data := d{
		Ir: a,
	}

	tmp.Execute(w, data)
}
