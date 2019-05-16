package vm

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error:", err)
		return false
	}

	return user.CheckPassword(password)
}
