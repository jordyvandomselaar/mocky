package serviceProviders

import "net/http"

// InitStaticFileServiceProvider inititializes the service provider for static files so the webserver returns static files when the url matches a file.
func InitStaticFileServiceProvider() {

	// Handlers for public files
	nodeModulesHandler := http.FileServer(http.Dir("./node_modules"))
	staticFileHandler := http.FileServer(http.Dir("./public"))

	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", nodeModulesHandler))
	http.Handle("/public/", http.StripPrefix("/public/", staticFileHandler))
}
