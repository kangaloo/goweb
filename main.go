package main

import (
	"github.com/kangaloo/goweb/controller"
	"log"
	"net/http"
)

func main() {
	controller.Startup()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
