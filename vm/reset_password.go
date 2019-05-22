package vm

import "github.com/kangaloo/goweb/model"

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

func GetResetPasswordViewModel(token string) *ResetPasswordViewModel {
	v := &ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
