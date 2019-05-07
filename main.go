package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

type Post struct {
	User User
	Body string
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	v := &IndexViewModel{
		Title: "home page",
		User: User{
			Username: "kangaroo",
		},
		Posts: []Post{
			{
				User: User{Username: "Alan"},
				Body: "hello world!",
			},
			{
				User: User{Username: "Alex"},
				Body: "hello world!",
			},
		},
	}

	tpl, err := template.New("index.html").ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("parse files error: %v", err)
	}
	/*
		https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
		2019/05/05 15:02:09 exec error template: "welcome" is an incomplete or empty template
	*/

	if err := tpl.Execute(w, v); err != nil {
		log.Fatalf("exec template error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", IndexHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
