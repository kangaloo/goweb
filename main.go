package main

import (
	"github.com/gorilla/context"
	"github.com/kangaloo/goweb/controller"
	"github.com/kangaloo/goweb/model"
	"log"
	"net/http"
)

func main() {
	db := model.ConnectToDB()
	//defer func() { _ = db.Close() }()
	defer db.Close()
	model.SetDB(db)
	controller.Startup()
	log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux)))
}
