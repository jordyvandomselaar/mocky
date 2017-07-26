package serviceProviders

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/routes"
	"net/http"
)

// InitRouteServiceProvider initializes the route service provider so Gorilla's Mux router will be used instead of the default router.
func InitRouteServiceProvider(router *mux.Router) {
	// Register handlers
	http.Handle("/", router)
}

// Initialize routes
func InitRoutes(router *mux.Router, templateManager managers.Template) {
	routes.InitHomeRoutes(router, templateManager)
	routes.InitAuthRoutes(router, templateManager)
}