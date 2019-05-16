package controller

import (
	"github.com/kangaloo/goweb/vm"
	"log"
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	user, _ := getSessionUser(r)
	log.Println(user)
	v := vm.GetVM(user)
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

		log.Println("login handler: will checkLogin")
		errs := checkLogin(username, password)
		v.AddError(errs...)

		if len(v.Errs) != 0 {
			_ = tpl.Execute(w, v)
		} else {
			_ = setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "register.html"
	v := vm.GetRegisterViewModel()

	if r.Method == http.MethodGet {
		_ = templates[tplName].Execute(w, v)
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			_ = templates[tplName].Execute(w, v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				_, _ = w.Write([]byte("Error insert database"))
				return
			}
			_ = setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}

}
