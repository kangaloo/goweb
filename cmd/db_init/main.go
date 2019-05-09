package main

import (
	"github.com/kangaloo/goweb/model"
	"log"
)

func main() {
	log.Printf("DB init...\n")

	db := model.ConnectToDB()
	defer func() { _ = db.Close() }()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})
}
