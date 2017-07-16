package helpers

import (
	"html/template"
	"net/http"
	"github.com/jordyvandomselaar/mock-backend/App/helpers/errors"
)

type View struct {
	Path string
}

func (view *View) Render(writer http.ResponseWriter, data interface{}) {
	tpl, err := template.ParseFiles("App/Views/layouts/app.gohtml", view.Path+".gohtml")
	errors.Handle(err)
	tpl.Execute(writer, data)
}
