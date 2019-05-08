package vm

import "github.com/kangaloo/goweb/model"

type BaseViewModel struct {
	Title string
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
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
	model.User
	Posts []model.Post
}

func GetVM() *IndexViewModel {

	u1 := model.User{Username: "Alan"}
	u2 := model.User{Username: "Alex"}

	posts := []model.Post{
		{User: u1, Body: "Beautiful day in Portland!"},
		{User: u2, Body: "The Avengers movie was so cool!"},
	}

	return &IndexViewModel{
		BaseViewModel{Title: "home page"},
		model.User{Username: "kangaroo"},
		posts,
	}
}
