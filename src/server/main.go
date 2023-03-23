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

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func setUpRoutes(r *mux.Router, db *gorm.DB) {
	log.Println("Setting up routes...")
	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignUp(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/signin", handlers.Login).Methods("POST")
	r.HandleFunc("/refresh", handlers.Refresh)
	r.HandleFunc("/logout", handlers.Logout)
	r.HandleFunc("/file/metadata/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Metadata(w, r, db)
	}).Methods("GET", "POST")

	log.Println("Routes set.")
}

func initServer() {
	log.Println("Vaulter is coming...")
	log.Println("Initializing server...")
	loadEnvVars()
	db := database.Connect()
	r := mux.NewRouter()
	http.Handle("/", r)
	setUpRoutes(r, db)
}

func main() {
	initServer()
	log.Println("Server is running on http://127.0.0.1:8888")
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
