package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/routes"
	"github.com/jordyvandomselaar/mock-backend/app/serviceProviders"
	"log"
	"net/http"
)

func main() {
	// GORM
	gormServiceProvider := serviceProviders.NewGormServiceProvider()
	defer gormServiceProvider.Db.Close()

	// Template Manager
	templateManager := managers.NewTemplateManager()

	// Router
	router := mux.NewRouter()

	routes.InitHomeRoutes(router, templateManager)
	serviceProviders.InitRouteServiceProvider(router)
	serviceProviders.InitStaticFileServiceProvider()

	// Fire up our server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
