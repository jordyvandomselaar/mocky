package helpers

import (
	"github.com/jordyvandomselaar/mock-backend/App/helpers/errors"
	"html/template"
	"net/http"
)

type View struct {
	Path string
}

func (view *View) Render(writer http.ResponseWriter, data interface{}) {
	tpl, err := template.ParseFiles("App/Views/layouts/app.gohtml", view.Path+".gohtml")
	errors.Handle(err)
	tpl.Execute(writer, data)
}
