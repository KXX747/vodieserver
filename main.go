package main

import (
	"net/http"
	"fmt"
	"github.com/KXX747/vodieserver/handler"
	_ "net/http/pprof"
)


func main() {

	fmt.Println("hello stream!")
	registHandler:=handler.RegisterHandlers()

	mh:=handler.NewMiddleHandler(registHandler)
	go func(mh http.Handler) {
		http.ListenAndServe(":8081",mh)
	}(mh)
	go func(mh http.Handler) {
		http.ListenAndServe(":8082",mh)
	}(mh)


	http.ListenAndServe(":8080",mh)

}

func GetUserName() string {

	return "2"

}


