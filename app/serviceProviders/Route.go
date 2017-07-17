package serviceProviders

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouteServiceProvider(router *mux.Router) {
	// Register handlers
	http.Handle("/", router)
}
