package forum

import (
	"html/template"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/CNA" {
		http.Error(w, "page not found", 404)
		return
	}

	tmp, err := template.ParseFiles("template/newAccount.html")

	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}

	tmp.Execute(w, nil)

}
