package vm

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}

func GetRegisterViewModel() *RegisterViewModel {
	v := &RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Printf("can not find user: %s \n", username)
		return false
	}
	return true
}

func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
