package serviceProviders

import (
	"github.com/gorilla/mux"
	"net/http"
)

// InitRouteServiceProvider initializes the route service provider so Gorilla's Mux router will be used instead of the default router.
func InitRouteServiceProvider(router *mux.Router) {
	// Register handlers
	http.Handle("/", router)
}
