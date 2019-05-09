package model

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `grom:"type:varchar(64)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Email        string `grom:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}
