package vm

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

type BaseViewModel struct {
	Title       string
	CurrentUser string
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
}

func GetVM(username string) *IndexViewModel {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println(err.Error())
		return &IndexViewModel{BaseViewModel{Title: "Home Page"}, nil}
	}

	posts, _ := model.GetPostsByUserID(user.ID)
	v := &IndexViewModel{BaseViewModel{Title: "Home Page"}, *posts}
	v.SetCurrentUser(username)
	return v
}
