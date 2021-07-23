package controllers

import (
	"net/http"
)

// 初期画面（ログイン不要）
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {

	// 	log.Fatalln(err)
	// }
	// t.Execute(w, "hello")
}

// ログインした人がいけるページ
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
