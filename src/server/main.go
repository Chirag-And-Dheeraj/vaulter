package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"strings"

	// "server/middlewares"
	"server/database"
)

func main() {

	loadEnvVars()
	database.Connect()
	http.HandleFunc("/signup", handlers.Signup)
	http.HandleFunc("/signin", handlers.Login)
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/logout", handlers.Logout)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func loadEnvVars() {
	log.Println("Setting environment variables...")

	envFile, err := os.Open("./.env")

	if err != nil {
		log.Fatal(err)
	}

	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), "=")
		os.Setenv(lineSplit[0], lineSplit[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Environment variables set.")
}
