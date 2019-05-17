package controller

import (
	"fmt"
	"github.com/gorilla/mux"
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

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	v, err := vm.GetProfileViewModel(sUser, pUser)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) does not exist", pUser)
		_, _ = w.Write([]byte(msg))
		return
	}
	_ = templates[tplName].Execute(w, &v)
}

func ProfileEditHandler(w http.ResponseWriter, r *http.Request) {
	tplName := "profile_edit.html"
	username, _ := getSessionUser(r)
	v := vm.GetProfileEditViewModel(username)

	if r.Method == http.MethodGet {
		if err := templates[tplName].Execute(w, v); err != nil {
			log.Println(err)
			return
		}
	}

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		aboutMe := r.Form.Get("aboutme")
		log.Println("get about me from user form post: ", aboutMe)
		if err := vm.UpdateAboutMe(username, aboutMe); err != nil {
			log.Println("update about me error: ", err)
			_, _ = w.Write([]byte("Error update about me"))
			return
		}
		http.Redirect(w, r, "", http.StatusSeeOther)
	}
}
