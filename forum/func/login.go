package forum

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/login" {
		http.Error(w, "page not found", 404)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", email).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		http.ServeFile(w, r, "login.html")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		http.ServeFile(w, r, "login.html")
		return
	}
}
