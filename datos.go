package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func datos_login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("email")
	pass := r.FormValue("pass")

	if name == "" || pass == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	name = select_user(name, pass)

	if name == "" {
		http.Redirect(w, r, "/login?d=fail", http.StatusFound)
		return
	}

	fmt.Println(name)
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "user", Value: name, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)

}

var code int

var (
	name      string
	addremail string
	pass      string
)

func handlerMail(w http.ResponseWriter, r *http.Request) {

	if name == "" {
		name = r.FormValue("name")
		addremail = r.FormValue("email")
		pass = r.FormValue("password")
	}

	/*
		if !select_view_exist_user(name) {
			http.Redirect(w, r, "/register?d=user", http.StatusFound)
			return
		} else if !select_view_exist_mail(addremail) {
			http.Redirect(w, r, "/register?d=mail", http.StatusFound)
			return
		}

	*/
	rand.Seed(time.Now().UnixNano())
	code = rand.Intn(10000-5000) + 3000

	z := r.FormValue("e")

	if z == "error" {
		w.Write([]byte("<script>alert('Codigo incorrecto')</script>"))
		http.Redirect(w, r, "/verify_mail", http.StatusFound)
		return
	}

	if !email(addremail, code) {
		http.Redirect(w, r, "/register?d=e_mail", http.StatusFound)
		return
	}

	tmp, err := template.ParseFiles("public/register/verify_code/index.html")
	if err != nil {
		http.Error(w, "El error es: "+err.Error(), http.StatusBadRequest)
		return
	}
	tmp.Execute(w, nil)
}

func dMail(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("code")
	i, err := strconv.Atoi(f)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if i == code {
		fmt.Println("el codigo es el mismo")
		http.Redirect(w, r, "/verifynew", http.StatusFound)
		return
	} else {
		fmt.Println("erro al codigo codigo NO es el mismo")
		http.Redirect(w, r, "/verify_mail?e=error", http.StatusFound)
		return
	}

}

func datos_register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(name, addremail, pass)
	value := insert_new_user(name, addremail, pass)

	if !value {
		http.Redirect(w, r, "/register?d=fail", http.StatusFound)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "user", Value: name, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)
}
