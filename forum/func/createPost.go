package forum

import (
	"html/template"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login-page", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/create-Post" {
		http.Error(w, "page not found", 404)
		return
	}

	tmp, err := template.ParseFiles("template/createpost.html")
	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}

	tmp.Execute(w, nil)
}
