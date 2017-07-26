package routes

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/mock-backend/app/controllers"
	"github.com/jordyvandomselaar/mock-backend/app/managers"
)

// InitAuthRoutes initializes the authentication routes
func InitAuthRoutes(router *mux.Router, templateManager managers.Template) {
	urlManager := managers.NewUrlManager()
	authController := controllers.NewAuthController(templateManager)

	// Define our routes
	router.HandleFunc(managers.ParseUrl(urlManager.Authenticate), authController.ShowAuthenticate)
}
