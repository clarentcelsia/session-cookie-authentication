package main

import (
	"cookie/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/refresh", handler.Refresh)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/hello", handler.Hello)
	//here we use defaultServeMux
	log.Fatal(http.ListenAndServe(":8080", nil))
}
