package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	defer db.Close()
	if err != nil {
		panic(err)
	}

	consulta := fmt.Sprintf("insert into users values ('%s','%s', '%s')", username, email, pass)
	fmt.Println("Fase: consulta" + consulta)
	_, err = db.Query(consulta)

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
	consulta_sql := fmt.Sprintf("select * from users where (username = '%s' or email = '%s') && contraseña = '%s'", username, username, pass)
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
	return v.user
}

type data struct {
	username string
	//image       byte
	description string
}

func select_user_view(username string) (string, string) {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}
	v := data{}
	consulta_sql := fmt.Sprintf("select username, descripcion from dates where username = '%s'", username)
	mysql, err := db.Query(consulta_sql)
	if err != nil {
		return "usuario no registrado", ""
	}
	for mysql.Next() {
		mysql.Scan(&v.username, &v.description)
	}
	return v.username, v.description

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
