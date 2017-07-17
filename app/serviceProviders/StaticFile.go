package serviceProviders

import "net/http"

func InitStaticFileServiceProvider() {

	// Handlers for public files
	nodeModulesHandler := http.FileServer(http.Dir("./node_modules"))
	staticFileHandler := http.FileServer(http.Dir("./public"))

	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", nodeModulesHandler))
	http.Handle("/public/", http.StripPrefix("/public/", staticFileHandler))
}
