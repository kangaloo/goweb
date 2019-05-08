package controller

import (
	"github.com/kangaloo/goweb/vm"
	"log"
	"net/http"
)

func IndexHandle(w http.ResponseWriter, _ *http.Request) {
	v := vm.GetVM()
	tpl := templates["index.html"]

	if err := tpl.Execute(w, v); err != nil {
		log.Fatalf("exec template error: %v", err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tplName := "login.html"

	v := &vm.LoginViewModel{}
	v.SetTitle("Login")
	tpl := templates[tplName]

	if r.Method == http.MethodGet {
		if err := tpl.Execute(w, v); err != nil {
			log.Fatalf("exec template error: %s", err.Error())
		}
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatalf("parse form failed: %s", err.Error())
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username too short")
		}

		if len(password) < 6 {
			v.AddError("password too short")
		}

		if !check(username, password) {
			v.AddError("username password not correct, please input again")
		}

		if len(v.Errs) != 0 {
			_ = tpl.Execute(w, v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func check(username, password string) bool {
	if username == "admin" && password == "abc123" {
		return true
	}
	return false
}
