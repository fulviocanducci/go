package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE sources(int id, text name)")
	if err != nil {
		panic(err)
	}
}
