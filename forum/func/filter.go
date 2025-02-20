package forum

import (
	"fmt"
	"html/template"
	"net/http"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")

	// // reg, err := regexp.Compile(category)
	// if err != nil {
	// 	http.Error(w, "page not found", 404)
	// 	return
	// }
	rows, err := db.Query("SELECT title, descriptions, topic FROM posts WHERE topic LIKE ?", "%"+category+"%")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "h", 500)
		return
	}
	defer rows.Close()
	type Filter struct {
		Title   string
		Descriptions string
		Topic   string
	}
	var arrFilter []Filter
	for rows.Next() {
		var newFilter Filter
		var title, descriptions, topic string
		er := rows.Scan(&title, &descriptions, &topic)

		if er != nil {
			http.Error(w, "htp", 500)
			return
		}
		newFilter.Title = title
		newFilter.Descriptions = descriptions
		newFilter.Topic = topic
		arrFilter = append(arrFilter, newFilter)
	}
	tmp, err:= template.ParseFiles("template/filter.html")
	if err != nil {
		http.Error(w, "ht", 500)
		return
	}
	 tmp.Execute(w, arrFilter)
	// if reg.MatchString("") {

	// }

	// rows, err := db.Query("SELECT title, comment, topic FROM posts;")
}
