package vm

import (
	"github.com/kangaloo/goweb/config"
	"github.com/kangaloo/goweb/model"
)

type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

func GetEmailViewModel(email string) *EmailViewModel {
	v := &EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken()
	v.Server = config.GetServerURL()
	return v
}
