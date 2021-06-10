package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	root     = "b12a7d9f181e5f"
	key      = "b906087e"
	host     = "us-cdbr-east-04.cleardb.com:3306"
	database = "heroku_47385c5b7a6b7fa"
)

func insert_new_user(username, email, pass string) string {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", root, key, host, database)

	db, err := sql.Open("mysql", link)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	consulta := fmt.Sprintf("insert into users values ('%s','%s', '%s')", username, email, pass)
	_, err = db.Query(consulta)

	if err != nil {
		return "Nombre de usuario o email existente, intente con otro"
	}
	return "Usuario registrado exitosamente"
}
