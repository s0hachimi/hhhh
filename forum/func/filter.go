package forum

import (
	"fmt"
	"net/http"
	"regexp"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")

	reg, err := regexp.Compile(category)
	if err != nil {
		http.Error(w, "page not found", 404)
		return
	}

	if reg.MatchString("") {

	}

	rows, err := db.Query("SELECT title, comment, topic FROM posts;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
		return
	}
	defer rows.Close()
}
