package controller

import (
	"errors"
	"fmt"
	"github.com/kangaloo/goweb/vm"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func populateTemplates() map[string]*template.Template {
	basePath := "templates"
	res := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + err.Error() + "'")
			}
			defer func() { _ = f.Close() }()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}

			tpl := template.Must(layout.Clone())
			_, err = tpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}

			res[fi.Name()] = tpl
		}()
	}

	return res
}

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("val: ", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}

	fmt.Println("Username: ", username)
	fmt.Println("sessionID: ", session.ID)
	fmt.Printf("%#v\n", session)
	fmt.Printf("%#v\n", store)

	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values["user"] = username
	return session.Save(r, w)
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}

func checkLogin(username, password string) []string {
	var errs []string
	if err := checkUserPassword(username, password); len(err) > 0 {
		errs = append(errs, err)
	}
	return errs
}

func checkUserPassword(username, password string) string {
	if !vm.CheckLogin(username, password) {
		return fmt.Sprintf("Username and password is not correct.")
	}
	return ""
}

func checkUserExist(username string) string {
	if vm.CheckUserExist(username) {
		return fmt.Sprintf("Username already exist, please choose another username")
	}
	return ""
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w._]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

func checkRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

func addUser(username, password, email string) error {
	return vm.AddUser(username, password, email)
}
