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

	username := r.FormValue("username")
	password := r.FormValue("password")

	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
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

	// sessionToken := generateSessionToken()
	// expiration := time.Now().Add(24 * time.Hour)

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "session_token",
	// 	Value:   sessionToken,
	// 	Expires: expiration,
	// })

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// func generateSessionToken() string {
// 	bytes := make([]byte, 16)
// 	rand.Read(bytes)
// 	return hex.EncodeToString(bytes)
// }
