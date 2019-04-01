package main

import (
	"net/http"
	"fmt"
	"github.com/KXX747/vodieserver/handler"
)


func main() {

	fmt.Println("hello stream!")
	registHandler:=handler.RegisterHandlers()

	mh:=handler.NewMiddleHandler(registHandler)
	http.ListenAndServe(":8080",mh)
}

func GetUserName() string {

	return "2"

}


