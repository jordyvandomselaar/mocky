package routes

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/mock-backend/app/controllers"
	"github.com/jordyvandomselaar/mock-backend/app/managers"
)

// InitHomeRoutes initializes the home routes
func InitHomeRoutes(router *mux.Router, templateManager managers.Template) {
	urlManager := managers.NewUrlManager()
	homeController := controllers.NewHomeController(templateManager)

	// Define our routes
	router.HandleFunc(managers.ParseUrl(urlManager.Home), homeController.Home)
	router.HandleFunc(managers.ParseUrl(urlManager.About), homeController.About)
}
