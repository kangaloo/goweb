package vm

import "github.com/kangaloo/goweb/model"

type ProfileViewModel struct {
	BaseViewModel
	BasePageViewModel
	Posts          []model.Post
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	ProfileUser    model.User
}

func GetProfileViewModel(sUser, pUser string, page, limit int) (*ProfileViewModel, error) {
	v := &ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}

	posts, total, _ := model.GetPostsByUserIDPageAndLimit(u.ID, page, limit)
	v.ProfileUser = *u
	v.Editable = sUser == pUser
	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}

	v.SetBasePageViewModel(total, page, limit)
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}

// Follow func : A follow B
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Follow(b)
}

// UnFollow func : A unfollow B
func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.UnFollow(b)
}
