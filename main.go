package main

import (
	_ "github.com/kangaloo/goweb/controller"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
