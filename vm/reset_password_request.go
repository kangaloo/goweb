package vm

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

func GetResetPasswordRequestViewModel() *ResetPasswordRequestViewModel {
	v := &ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

func CheckEmailExist(email string) bool {
	if _, err := model.GetUserByEmail(email); err != nil {
		log.Println("Can not find email:", email)
		return false
	}
	return true
}
