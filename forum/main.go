package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	forum "forum/func"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./databases/data.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStatements, err := os.ReadFile("./databases/my.sql")
	if err != nil {
		log.Fatal("Error reading SQL file:", err)
	}
	_, err = db.Exec(string(sqlStatements))
	if err != nil {
		log.Fatal("Error executing SQL statements:", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	forum.InitHandlers(db)
}

func main() {
	defer db.Close()
	http.HandleFunc("/", forum.HomePage)
	http.HandleFunc("/login", forum.LoginPage)
	http.HandleFunc("/singup", forum.Singup)
	http.HandleFunc("/check", forum.Check)
	http.HandleFunc("/log", forum.Log)


	fmt.Println("http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
