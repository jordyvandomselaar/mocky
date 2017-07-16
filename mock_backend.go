package main

import (
	"github.com/jordyvandomselaar/mock-backend/App/controllers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	homeController := controllers.HomeController{}
	router := mux.NewRouter()
	nodeModulesHandler := http.FileServer(http.Dir("./"))

	router.HandleFunc("/", homeController.Home)

	http.Handle("/", router)
	http.Handle("/node_modules/", nodeModulesHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
