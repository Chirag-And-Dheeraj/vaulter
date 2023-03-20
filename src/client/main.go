package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setUpRoutes(r *mux.Router) {
	log.Println("Setting up routes...")
	r.HandleFunc("/", LandingHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/setup", SetupHandler).Methods("GET")
	r.HandleFunc("/vault", VaultHandler).Methods("GET")
	r.HandleFunc("/viewer", ViewerHandler).Methods("GET")
	r.HandleFunc("/dashboard", DashboardHandler).Methods("GET")
	http.Handle("/", r)
	log.Println("Routes set.")
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/index.html"
		http.ServeFile(w, r, p)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/login.html"
		http.ServeFile(w, r, p)
	}
}

func SetupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/setup" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/setup.html"
		http.ServeFile(w, r, p)
	}
}

func VaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/vault" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/vault.html"
		http.ServeFile(w, r, p)
	}
}

func ViewerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/viewer" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/viewer.html"
		http.ServeFile(w, r, p)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard" {
		return
	} else if r.Method == "GET" {
		log.Println("GET: " + r.URL.Path)
		p := "./pages/dashboard.html"
		http.ServeFile(w, r, p)
	}
}

func initServer() {
	log.Println("Vaulter is coming...")
	log.Println("Initializing server...")
	// utils.LoadEnvVars()
	// db := database.Connect()
	r := mux.NewRouter()
	setUpRoutes(r)
}

func main() {
	initServer()
	log.Println("Server is running on http://127.0.0.1:7777")
	log.Fatal(http.ListenAndServe("127.0.0.1:7777", nil))
}
