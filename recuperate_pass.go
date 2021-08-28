package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func recuperate(w http.ResponseWriter, r *http.Request) {

	e := r.FormValue("e")

	if e == "error" {
		w.Write([]byte("<alert>Error al verificar email</alert>"))
	}

	t, err := template.ParseFiles("public/recuperate_pass/pass.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

var (
	mail string
)

func handleRecuperate(w http.ResponseWriter, r *http.Request) {
	mail = r.FormValue("mail")

	if mail == "" {
		http.Redirect(w, r, "/pass", http.StatusFound)
		return
	}
	rand.Seed(time.Now().UnixNano())
	code = rand.Intn(10000-5000) + 3000

	enviadoMail := recuperate_password_email(mail, code)

	if !enviadoMail {
		http.Redirect(w, r, "/pass?e=error", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/verify_code", http.StatusFound)

}

//Diseños

func handleVerifyCode(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/recuperate_pass/change/index.html")

	if err != nil {
		http.Error(w, "el error es: "+err.Error(), http.StatusBadRequest)
	}

	t.Execute(w, nil)
}

func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/recuperate_pass/change/change.html")

	if err != nil {
		http.Error(w, "el error es"+err.Error(), http.StatusBadRequest)
	}

	t.Execute(w, nil)
}

//Backend
func verify_code(w http.ResponseWriter, r *http.Request) {
	code_email := r.FormValue("code")
	a, _ := strconv.Atoi(code_email)
	if a == code {
		http.Redirect(w, r, "/change_password", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/verify_code", http.StatusFound)
}

func change_password(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<script>alert('Todo disponible para cambiar la contraseña')</script>"))
}
