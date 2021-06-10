package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "us-cdbr-east-03.cleardb.com"
	user     = "bc93937ad47f25"
	pass     = "10174fc1"
	database = "heroku_81792ca3b705421"
)

type pagina struct {
	title string
	main  string
	url   string
}

func main() {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database)
	db, err := sql.Open("mysql", link)

	if err != nil {
		panic(err)
	}

	v := pagina{}
	mysql, err := db.Query("select * from html")
	if err != nil {
		panic(err)
	}
	for mysql.Next() {
		err = mysql.Scan(&v.title, &v.main, &v.url)
	}
	if err != nil {
		fmt.Println("No mostro usuarios")
	}
	fmt.Println(v.title)
}
