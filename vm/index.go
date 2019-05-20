package vm

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

type IndexViewModel struct {
	BaseViewModel
	BasePageViewModel
	Posts []model.Post // current user's following posts
	Flash string
}

func GetVM(username, flash string, page, limit int) *IndexViewModel {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println(err.Error())
		return &IndexViewModel{BaseViewModel: BaseViewModel{Title: "Home Page"}, BasePageViewModel: BasePageViewModel{}, Flash: flash}
	}

	posts, total, err := user.FollowingPostsByPageAndLimit(page, limit)
	if err != nil {
		log.Println("vm.GetVM user.FollowingPostsByPageAndLimit failed.")
		return &IndexViewModel{BaseViewModel: BaseViewModel{Title: "Home Page"}, BasePageViewModel: BasePageViewModel{}, Flash: flash}
	}

	v := &IndexViewModel{}
	v.SetTitle("Home Page")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, post string) error {
	u, err := model.GetUserByUsername(username)
	if err != nil {
		return err
	}
	return u.CreatePost(post)
}
