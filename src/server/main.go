package main

import (
	"log"
	"net/http"
	"server/handlers"
)


func main() {

	LoadEnvVars()

	

	http.HandleFunc("/signin", handlers.Login)
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/logout", handlers.Logout)

	log.Fatal(http.ListenAndServe(":8888", nil))
}