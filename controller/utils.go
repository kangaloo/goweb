package controller

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/kangaloo/goweb/config"
	"github.com/kangaloo/goweb/vm"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
	//if m, _ := regexp.MatchString(`^([\w._]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
	//	return fmt.Sprintf("Email field not a valid email")
	//}
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

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	_ = session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}

	_ = session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func getPage(r *http.Request) int {
	url := r.URL
	query := url.Query()
	p := query.Get("page")
	if p == "" {
		return 1
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		return 1
	}
	return page
}

func sendEmail(target, subject, content string) {
	server, port, user, pwd := config.GetSMTPConfig()
	d := gomail.NewPlainDialer(server, port, user, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", target)
	m.SetAddressHeader("Cc", user, "admin")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		log.Println("sendEmail: ", err)
		return
	}

	log.Printf("sendMail: send mail to %s successfully\n", target)
}

func checkResetPasswordRequest(email string) []string {
	var errs []string
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmailExist(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

func checkEmailExist(email string) string {
	if !vm.CheckEmailExist(email) {
		return fmt.Sprintf("Email does not register yet.Please Check email.")
	}
	return ""
}

func checkResetPassword(pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	return errs
}
