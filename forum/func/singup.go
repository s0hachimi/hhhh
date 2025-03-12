package forum

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func InitHandlers(database *sql.DB) {
	db = database
}

type signupReq struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password1 string `json:"pass1"`
	Password2 string `json:"pass2"`
}

func Singup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	var req signupReq

	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	hh := Hh(req)
	if hh != "" {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": hh,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password1), bcrypt.DefaultCost)
	if err != nil {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?);", req.Username, req.Email, string(hashedPassword))
	if err != nil {
		fmt.Printf("%#v\n", err.Error()[26:])
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": err.Error()[26:],
		})
		return
	}

	sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
		"success": true,
	})
}

func Hh(req signupReq) string {

	if !isGmail(req.Email) {
		return "Email must be a valid Gmail address."
	}

	if req.Password1 != req.Password2 {
		return "Password not match"
	}

	if !isStrongPassword(req.Password1) {
		return "Password must be at least 8 characters long and include at least one number, one lowercase letter, and one uppercase letter."
	}
	
	return ""
}

func isGmail(email string) bool {
	// match, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9._%+-]*@gmail\\.com$", email)
	gmail, _ := regexp.MatchString("^[a-zA-Z]", email)
	end, _ := regexp.MatchString("@gmail.com$", email)
	return gmail && end
}

func isStrongPassword(password string) bool {
    hasLower, _ := regexp.MatchString("[a-z]", password)
    hasUpper, _ := regexp.MatchString("[A-Z]", password)
    hasNumber, _ := regexp.MatchString("[0-9]", password)
    isLong := len(password) >= 8

    return hasLower && hasUpper && hasNumber && isLong
}
