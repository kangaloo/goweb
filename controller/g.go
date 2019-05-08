package controller

import (
	"html/template"
	"net/http"
)

var (
	templates map[string]*template.Template
)

func init() {
	templates = populateTemplates()
}

func Startup() {
	registerRoutes()
}

func registerRoutes() {
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/login", LoginHandler)
}
