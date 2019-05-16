package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

var (
	templates   map[string]*template.Template
	sessionName string
	store       *sessions.CookieStore
)

func init() {
	templates = populateTemplates()
	sessionName = "go-web"
	store = sessions.NewCookieStore([]byte("go-web-secret"))
}

func Startup() {
	registerRoutes()
}

func registerRoutes() {
	http.HandleFunc("/", middleAuth(IndexHandle))
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", middleAuth(LogoutHandler))
	http.HandleFunc("/register", RegisterHandler)
}
