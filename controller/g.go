package controller

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	r.HandleFunc("/", middleLog(middleAuth(indexHandle)))
	r.HandleFunc("/login", middleLog(loginHandler))
	r.HandleFunc("/logout", middleLog(middleAuth(logoutHandler)))
	r.HandleFunc("/register", middleLog(registerHandler))
	r.HandleFunc("/user/{username}", middleLog(middleAuth(profileHandler)))
	r.HandleFunc("/follow/{username}", middleLog(middleAuth(followHandler)))
	r.HandleFunc("/unfollow/{username}", middleLog(middleAuth(unFollowHandler)))
	r.HandleFunc("/profile_edit", middleLog(middleAuth(profileEditHandler)))
	r.HandleFunc("/explore", middleLog(middleAuth(exploreHandler)))
	r.HandleFunc("/reset_password_request", middleLog(resetPasswordRequestHandler))
	r.HandleFunc("/reset_password/{token}", middleLog(resetPasswordHandler))
	http.Handle("/", r)
}
