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
	flashName   string
	store       *sessions.CookieStore
	pageLimit   int
)

func init() {
	templates = populateTemplates()
	sessionName = "go-web"
	flashName = "go-web-flash"
	store = sessions.NewCookieStore([]byte("go-web-secret"))
	pageLimit = 5
}

func Startup() {
	registerRoutes()
}

func registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleAuth(indexHandle))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unFollowHandler))
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/explore", middleAuth(exploreHandler))
	http.Handle("/", r)
}
