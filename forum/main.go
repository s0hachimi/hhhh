package main

import (
	forum "forum/func"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", forum.HomePage)
	http.HandleFunc("/login", forum.LoginPage)
	http.HandleFunc("/CNA", forum.CreateAccount)
	http.HandleFunc("/check", forum.DataBase)

	

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
