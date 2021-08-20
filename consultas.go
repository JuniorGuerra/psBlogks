package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	root     = "b12a7d9f181e5f"
	key      = "b906087e"
	host     = "us-cdbr-east-04.cleardb.com:3306"
	database = "heroku_47385c5b7a6b7fa"
)

func insert_new_user(username, email, pass string) bool {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)

	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	passByte := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return err != nil
	}
	hashStringPass := string(hash)

	consulta := fmt.Sprintf("insert into users values ('%s','%s', '%s')", username, email, hashStringPass)

	fmt.Println("Fase: consulta" + consulta)
	_, err = db.Query(consulta)

	update_data_profile("", username, "", "")

	db.Close()
	return err == nil
}

type user_registred struct {
	user  string
	email string
	pass  string
}

func select_user(username, pass string) string {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := user_registred{}
	consulta_sql := fmt.Sprintf("select * from users where (username = '%s' or email = '%s')", username, username)

	mysql, err := db.Query(consulta_sql)
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		err = mysql.Scan(&v.user, &v.email, &v.pass)
	}

	if err != nil {
		return ""
	}

	PassHashByte := []byte(v.pass)
	passByte := []byte(pass)
	verify := bcrypt.CompareHashAndPassword(PassHashByte, passByte)

	if verify != nil {
		fmt.Println("Las contraseñas NOPE coinciden")
		return ""
	}

	fmt.Println(v.user, v.email)
	return v.user
}

//Sucia recoleccion del gmail, busqueda mas segura proximamente cuando busque mysql worbeach ahhaha

func select_gmail(username string) string {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := user_registred{}
	consulta_sql := fmt.Sprintf("select * from users where username = '%s'", username)
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		err = mysql.Scan(&v.user, &v.email, &v.pass)
	}
	if err != nil {
		return ""
	}
	fmt.Println(v.user, v.email)
	return v.email
}

type data struct {
	username    string
	image       string
	phone       string
	description string
}

func select_user_view(username string) (string, string, string, string) {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}
	v := data{}
	consulta_sql := fmt.Sprintf("select * from dates where username = '%s'", username)
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		fmt.Println(err)
		return "usuario no registrado", "", "", ""
	}
	for mysql.Next() {
		mysql.Scan(&v.image, &v.username, &v.phone, &v.description)
	}
	return v.image, v.username, v.phone, v.description

}

func insert_new_book(title, body, author, category string) string {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	id := GenerateToken(title)
	t := time.Now()
	fecha := t.Format("2006-01-02 15:04:05")

	consulta_sql := fmt.Sprintf("insert into post values ('%s', '%s', '%s','%s','%s', '%s' )", id, title, body, author, category, fecha)

	_, err = db.Query(consulta_sql)
	if err != nil {
		return "No se ha podido publicar el libro"
	}
	return "Libro publicado correctamente"
}

func update_data_profile(img, name, phone, description string) {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	consulta_sql := fmt.Sprintf("update dates set image_profile = '%s', telefono = '%s', descripcion = '%s' where username = '%s'", img, phone, description, name)

	_, err = db.Query(consulta_sql)

	if err != nil {
		panic(err)
	}
}

type users_registred struct {
	name string
	info string
}

func select_users_query_all(date string) []users_registred {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := users_registred{}
	va := []users_registred{}
	consulta_sql := "select users.username, dates.descripcion from users, dates where users.username = dates.username && users.username like " + "'%" + date + "%'"
	fmt.Println(consulta_sql)
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		mysql.Scan(&v.name, &v.info)
		va = append(va, v)
	}
	fmt.Println("Consulta", va)
	return va
}
