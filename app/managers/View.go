package managers

import (
	"html/template"
)

// View pre-compiles all views and makes sure all views all valid.
type View struct {
	Home         *template.Template
	Authenticate *template.Template
	About        *template.Template
}

// NewViewManager returns a new View manager.
func NewViewManager() *View {
	return &View{
		Home:         ParseView("home/index"),
		About:        ParseView("home/about"),
		Authenticate: ParseView("auth/authenticate"),
	}
}

// ParseView compiles a view and panics if the view contains an error.
func ParseView(name string) *template.Template {
	name = "app/views/" + name + ".gohtml"
	layoutPath := "app/views/layouts/app.gohtml"

	return template.Must(template.ParseFiles(layoutPath, name))
}
