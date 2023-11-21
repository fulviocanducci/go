package program

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE TABLE sources(int id, text name)")
}
