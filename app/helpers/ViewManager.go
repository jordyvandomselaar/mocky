package helpers

import (
	"html/template"
)

type ViewManager struct {
	Home         *template.Template
	Authenticate *template.Template
}

//func (view *View) Render(writer http.ResponseWriter, data interface{}) {
//	tpl, err := template.ParseFiles("app/Views/layouts/app.gohtml", "app/views/"+view.Path+".gohtml")
//	errors.Handle(err)
//	tpl.Execute(os.Stdout, data)
//}

func NewViewManager() *ViewManager {
	return &ViewManager{
		Home:         MustParseView("home/index"),
		Authenticate: MustParseView("auth/authenticate"),
	}
}

func MustParseView(name string) *template.Template {
	name = "app/views/" + name + ".gohtml"
	layoutPath := "app/views/layouts/app.gohtml"

	return template.Must(template.ParseFiles(layoutPath, name))
}
