package main

import (
	"fmt"
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

	if r.FormValue("e") == "error" {
		w.Write([]byte("<script>alert('Codigo incorrecto')</script>"))
	}

	if err != nil {
		http.Error(w, "el error es: "+err.Error(), http.StatusBadRequest)
	}

	t.Execute(w, nil)
}

func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/recuperate_pass/change/change.html")

	e := r.FormValue("e")

	if e == "vacio" {
		w.Write([]byte("<script>alert('Los datos de contraseñas no deben ser vacios')</script>"))
	}
	if e == "error" {
		w.Write([]byte("<script>alert('Las contraseñas no coinciden')</script>"))
	}
	if e == "fail" {
		w.Write([]byte("<script>alert('Un error al cambiar la contraseña, comuniquese a soporte')</script>"))
	}
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
	http.Redirect(w, r, "/verify_code?e=error", http.StatusFound)
}

func change_password(w http.ResponseWriter, r *http.Request) {
	var pass string = r.FormValue("pass")
	var verify_pass string = r.FormValue("verify_pass")

	if pass == "" || verify_pass == "" {
		w.Write([]byte("<script>alert('Los datos de contraseñas no deben ser vacios')</script>"))
		http.Redirect(w, r, "/change_password?e=vacio", http.StatusFound)
		return

	} else if pass != verify_pass {
		http.Redirect(w, r, "/change_password?e=error", http.StatusFound)
		return
	}

	verificacion := query_change_password(pass, mail)
	fmt.Println(pass, mail)

	if !verificacion {
		http.Redirect(w, r, "/change_password?e=fail", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/login?d=exitoso", http.StatusFound)
}
