package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func DataBase(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT,
    password TEXT
)
`)
	if err != nil {
		fmt.Println(err)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("passA")

	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		fmt.Println(err)
		return
	}

}
