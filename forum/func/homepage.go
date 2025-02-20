package forum

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "page not found", 404)
		return
	}

	tmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}

	rows, err := db.Query("SELECT title, descriptions, time, topic FROM posts ORDER BY time DESC;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
		return
	}
	defer rows.Close()

	type Post struct {
		Title   string
		Descriptions string
		Time string
		Topic   string
	}

	var arrPost []Post

	for rows.Next() {
		var newPost Post
		var title, descriptions,t, topic string
		er := rows.Scan(&title, &descriptions,&t, &topic)

		if er != nil {
			http.Error(w, "htppp", 500)
			return
		}

		newPost.Title = title
		newPost.Descriptions = descriptions
		newPost.Time = t
		newPost.Topic = topic
		arrPost = append(arrPost, newPost)

	}
	err = tmp.Execute(w, arrPost)

    if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
			return
	}
}
