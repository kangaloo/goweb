package controller

import (
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.HandleFunc("/", middleAuth(IndexHandle))
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", middleAuth(LogoutHandler))
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/user/{username}", middleAuth(ProfileHandler))
	r.HandleFunc("/profile_edit", middleAuth(ProfileEditHandler))
	http.Handle("/", r)
}
