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


	rows, err := db.Query("SELECT id, username, title, descriptions, time, topic, likes, dislikes FROM posts ORDER BY time DESC;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
		return
	}
	defer rows.Close()

	type users struct {
		IsLoggedIn bool
		Username   string
	}

	type Post struct {
		ID           int
		Username     string
		Title        string
		Descriptions string
		Time         string
		Topic        string
		Likes        int
		Dislikes     int
		User         users
	}

	var arrPost []Post
	var newPost Post

	is, username := IsLoggedIn(r)
	newPost.User.Username = username
	newPost.User.IsLoggedIn = is

	arrPost = append(arrPost, newPost)

	for rows.Next() {
		var username, title, descriptions, t, topic string
		var id, like, dislike int
		er := rows.Scan(&id, &username, &title, &descriptions, &t, &topic, &like, &dislike)

		if er != nil {
			fmt.Println(er)
			http.Error(w, "database !", 500)
			return
		}

		newPost.ID = id
		newPost.Username = username
		newPost.Title = title
		newPost.Descriptions = descriptions
		newPost.Time = t
		newPost.Topic = topic
		newPost.Likes = like
		newPost.Dislikes = dislike

		arrPost = append(arrPost, newPost)
		newPost = Post{}
	}


	tmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}

	err = tmp.Execute(w, arrPost)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "htppp", 500)
		return
	}
}
