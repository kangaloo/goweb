package main

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

func main() {
	dbInit()
}

func dbInit() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	_ = model.AddUser("admin", "abc123", "admin@139.com")
	_ = model.AddUser("kangaroo", "abc123", "kangaroo@139.com")

	u1, _ := model.GetUserByUsername("admin")
	_ = u1.CreatePost("Beautiful day in Portland!")
	_ = model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("kangaroo")
	_ = u2.CreatePost("The Avengers movie was so cool!")
	_ = u2.CreatePost("Sun shine is beautiful")

	_ = u1.Follow(u2.Username)
}
