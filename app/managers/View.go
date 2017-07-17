package managers

import (
	"html/template"
)

type View struct {
	Home         *template.Template
	Authenticate *template.Template
	About        *template.Template
}

func NewViewManager() *View {
	return &View{
		Home:         MustParseView("home/index"),
		About:        MustParseView("home/about"),
		Authenticate: MustParseView("auth/authenticate"),
	}
}

func MustParseView(name string) *template.Template {
	name = "app/views/" + name + ".gohtml"
	layoutPath := "app/views/layouts/app.gohtml"

	return template.Must(template.ParseFiles(layoutPath, name))
}
