package forum

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func InitHandlers(database *sql.DB) {
	db = database
}


func Check(w http.ResponseWriter, r *http.Request) {
	
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("pass2")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to hash password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?);", username, email, string(hashedPassword))
	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "template/emailCheck.html")
		return
	}

	http.ServeFile(w, r, "template/index.html")
}
