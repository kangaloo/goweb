package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

type Post struct {
	User
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

	tpl := PopulateTemplates()["index.html"]

	if err := tpl.Execute(w, v); err != nil {
		log.Fatalf("exec template error: %v", err)
	}
}

func PopulateTemplates() map[string]*template.Template {
	basePath := "templates"
	res := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + err.Error() + "'")
			}
			defer func() { _ = f.Close() }()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}

			tpl := template.Must(layout.Clone())
			_, err = tpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}

			res[fi.Name()] = tpl
		}()
	}

	return res
}

func main() {
	http.HandleFunc("/", IndexHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
