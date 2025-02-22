package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/login" {
		http.Error(w, "page not found", 404)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var userID int
	var hashedPassword string
	err := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &hashedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "username is incorrect !", 404)
		return
	} else if err != nil {
		http.Error(w, "Error in Database !", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		http.Error(w, "password is incorrect !", 404)
		return
	}

	sessionToken := uuid.New().String()
	expiration := time.Now().Add(24 * time.Hour)

	_, err = db.Exec("UPDATE users SET session_token = ? WHERE id = ?", sessionToken, userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error saving session!", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func IsLoggedIn(r *http.Request) (bool, string) {
	
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return false, ""
	}

	var userName string
	err = db.QueryRow("SELECT username FROM users WHERE session_token = ?", cookie.Value).Scan(&userName)
	if err == sql.ErrNoRows || err != nil {
		return false, ""
	}

	return true, userName
}

