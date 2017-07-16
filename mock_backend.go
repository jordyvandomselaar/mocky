package main

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/mock-backend/app/controllers"
	"log"
	"net/http"
)

func main() {
	// Controllers
	homeController := controllers.HomeController{}
	authController := controllers.AuthController{}

	// Router
	router := mux.NewRouter()

	// Handlers for public files
	nodeModulesHandler := http.FileServer(http.Dir("./node_modules"))
	staticFileHandler := http.FileServer(http.Dir("./public"))

	// Define our routes
	router.HandleFunc("/", homeController.Home)
	router.HandleFunc("/auth/authenticate", authController.ShowLogin)

	// Register handlers
	http.Handle("/", router)
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", nodeModulesHandler))
	http.Handle("/public/", http.StripPrefix("/public/", staticFileHandler))

	// Fire up our server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
