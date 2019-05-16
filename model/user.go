package model

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `grom:"type:varchar(64)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Email        string `grom:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CheckPassword(password string) bool {
	return u.PasswordHash == GeneratePasswordHash(password)
}

func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

func AddUser(username, password, email string) error {
	user := &User{
		Username: username,
		Email:    email,
	}
	user.SetPassword(password)
	return db.Create(user).Error
}
