package vm

import "github.com/kangaloo/goweb/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	ProfileUser model.User
}

func GetProfileViewModel(sUser, pUser string) (*ProfileViewModel, error) {
	v := &ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}

	posts, _ := model.GetPostsByUserID(u.ID)
	v.ProfileUser = *u
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}
