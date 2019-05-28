package controller

import (
	"github.com/kangaloo/goweb/model"
	"log"
	"net/http"
	"time"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username, err := getSessionUser(request)
		log.Println("middle: ", username)

		if username != "" {
			log.Println("Update last seen: ", username)
			_ = model.UpdateLastSeen(username)
		}

		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(writer, request)
		}
	}
}

func middleLog(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		next(writer, request)
		log.Printf("%s %s %s done in %v", request.RemoteAddr, request.Method, request.URL.Path, time.Since(now))
	}
}
