package forum

import (
	"fmt"
	"net/http"
	"time"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login-page", http.StatusSeeOther)
		return
	}
	

	title := r.FormValue("title")
	description := r.FormValue("description")
	s := r.Form["topic"]
	hh := time.Now()
	t := hh.Format("2006-01-02 15:04:05")

	topic := ""

	for _, v := range s {
		topic += v + ","
	}

	_, err = db.Exec("INSERT INTO posts (title, descriptions, time, topic) VALUES (?, ?, ?, ?);", title, description, t, topic)
	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "template/emailCheck.html")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
