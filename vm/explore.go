package vm

import "github.com/kangaloo/goweb/model"

type ExploreViewModel struct {
	BaseViewModel
	BasePageViewModel
	Posts []model.Post
}

func GetExploreViewModel(username string, page, limit int) *ExploreViewModel {
	posts, total, _ := model.GetPostByPageAngLimit(page, limit)
	v := &ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}
