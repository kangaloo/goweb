package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := &User{
			Username: "kangaroo",
		}

		tpl, err := template.New("index.html").ParseFiles("templates/index.html")
		if err != nil {
			log.Fatalln("parse error", err)
		}

		/*
			https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
			2019/05/05 15:02:09 exec error template: "welcome" is an incomplete or empty template
		*/
		if err := tpl.Execute(w, user); err != nil {
			log.Fatalln("exec error", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
