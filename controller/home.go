package controller

import (
	"fmt"
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
	v := &vm.LoginViewModel{}
	v.SetTitle("Login")
	tpl := templates["login.html"]

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
		_, _ = fmt.Fprintf(w, "%s %s", username, password)
	}
}
