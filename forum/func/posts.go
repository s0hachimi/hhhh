package forum

import (
	"fmt"
	"net/http"
)


func Posts(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	comment := r.FormValue("comment")
	s := r.Form["topic"]

	topic := ""

	for _, v := range s {
		topic += v + ","
	}

	_, err := db.Exec("INSERT INTO posts (title, comment, topic) VALUES (?, ?, ?);", title, comment, topic)

	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "template/emailCheck.html")
		return
	}


	http.Redirect(w, r, "/", http.StatusSeeOther)
}
