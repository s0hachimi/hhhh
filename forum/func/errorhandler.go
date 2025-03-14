package forum

import (
	"html/template"
	"net/http"
)

type Error struct {
	Message string
	Code int
}

func ErrorHandler(w http.ResponseWriter, message string, code int) {

	var Er Error

	Er.Message = message
	Er.Code = code

	w.WriteHeader(code)

	tmp, err := template.ParseFiles("template/error.html")
	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}

	err = tmp.Execute(w, Er)
	if err != nil {
		http.Error(w, "htppp", 500)
		return
	}
}
