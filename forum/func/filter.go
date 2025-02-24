package forum

import (
	"fmt"
	"html/template"
	"net/http"
)

func Filter(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/filter" {
		http.Error(w, "page not found", 404)
		return
	}

	category := r.FormValue("category")


	rows, err := db.Query("SELECT id, title, descriptions, time, topic, likes, dislikes FROM posts WHERE topic LIKE ?", "%"+category+"%")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "h", 500)
		return
	}
	defer rows.Close()
	type Filter struct {
		ID int
		Title   string
		Descriptions string
		Time string
		Topic   string
		Likes int
		Dislikes int
		IsLoggedIn bool
		Username string
	}
	var arrFilter []Filter
	for rows.Next() {
		var newFilter Filter
		var title, descriptions, t, topic string
		var id, like, dislike  int
		er := rows.Scan(&id, &title, &descriptions, &t, &topic, &like, &dislike)

		if er != nil {
			http.Error(w, "htp", 500)
			return
		}
		newFilter.ID = id
		newFilter.Title = title
		newFilter.Descriptions = descriptions
		newFilter.Time = t
		newFilter.Topic = topic
		newFilter.Likes = like
		newFilter.Dislikes = dislike

		is, username := IsLoggedIn(r)
		newFilter.Username = username
		newFilter.IsLoggedIn = is

		arrFilter = append(arrFilter, newFilter)
	}
	tmp, err:= template.ParseFiles("template/filter.html")
	if err != nil {
		http.Error(w, "ht", 500)
		return
	}
	err = tmp.Execute(w, arrFilter)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
			return
	}
}
