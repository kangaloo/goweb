package controller

import (
	"html/template"
	"net/http"
)

var (
	templates map[string]*template.Template
)

func init() {
	templates = PopulateTemplates()
	registerRoutes()
}

func registerRoutes() {
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/login", LoginHandler)
}
