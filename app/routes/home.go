package routes

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/mock-backend/app/controllers"
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"strings"
)

// InitHomeRoutes initializes the home routes
func InitHomeRoutes(router *mux.Router, templateManager managers.Template) {
	urlManager := managers.NewUrlManager()
	homeController := controllers.NewHomeController(templateManager)
	authController := controllers.NewAuthController(templateManager)

	// Define our routes
	router.HandleFunc(getUrl(urlManager.Home), homeController.Home)
	router.HandleFunc(getUrl(urlManager.About), homeController.About)
	router.HandleFunc(getUrl(urlManager.Authenticate), authController.ShowAuthenticate)

}

func getUrl(url string) string {
	return strings.Split(url, "#")[0]
}
