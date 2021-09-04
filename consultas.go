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

/*
func select_view_exist_user(name string) bool {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)

	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	consulta := fmt.Sprintf("select username from users where username = %s ", name)

	fmt.Println("Fase: consulta" + consulta)
	mysql, err := db.Query(consulta)

	if err != nil {
		panic(err)
	}
	a := ""

	mysql.Scan(a)

	db.Close()

	return a == ""

}

func select_view_exist_mail(mail string) bool {

	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)

	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	consulta := fmt.Sprintf("select username from users where email = %s ", mail)
	mysql, err := db.Query(consulta)

	if err != nil {
		panic(err)
	}
	a := ""

	mysql.Scan(a)

	db.Close()

	return a == ""
}
*/
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

	insert_data_profile("", username, "", "")

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

type booksUser struct {
	id        string
	title     string
	body      string
	autor     string
	categoria string
	fecha     string
}

func selectAllBookUser(username string) []booksUser {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := booksUser{}
	va := []booksUser{}
	consulta_sql := "select * from post where autor = '" + username + "'"
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		mysql.Scan(&v.id, &v.title, &v.body, &v.autor, &v.categoria, &v.fecha)
		va = append(va, v)
	}
	db.Close()
	fmt.Println("Consulta", va)
	return va
}

func selectBookUser(username, title string) booksUser {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := booksUser{}
	consulta_sql := "select * from post where autor = '" + username + "' && titulo = '" + title + "'"
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		mysql.Scan(&v.id, &v.title, &v.body, &v.autor, &v.categoria, &v.fecha)
	}
	db.Close()
	return v
}

func insert_data_profile(img, name, phone, description string) {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	consulta_sql := fmt.Sprintf("insert into dates values('%s','%s','%s','%s')", img, name, phone, description)

	_, err = db.Query(consulta_sql)

	if err != nil {
		panic(err)
	}
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

func query_change_password(pass, name string) bool {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		fmt.Println(err.Error())
	}
	passByte := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return err != nil
	}
	hashStringPass := string(hash)

	consulta_sql := fmt.Sprintf("update users set contraseña = '%s' where email = '%s'", hashStringPass, name)

	_, err = db.Query(consulta_sql)

	db.Close()

	return err == nil
}
