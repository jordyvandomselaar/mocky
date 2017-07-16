package main

import (
	"log"
	"net/http"
	"github.com/jordyvandomselaar/mock-backend/App/controllers"
)

func main() {
	homeController := controllers.HomeController{}

	http.HandleFunc("/", homeController.Home)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


