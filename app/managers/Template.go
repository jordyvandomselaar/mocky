package managers

import (
	"html/template"
	"io"
	"log"
)

// Template pre-compiles all templates and makes sure all templates all valid.
type Template struct {
	Templates map[string]*template.Template
}

// Render renders a template.
func (tm Template) Render(template string, w io.Writer, data interface{}) {
	if tm.Templates[template] == nil {
		log.Fatal("Template not found: " + template)
	}

	err := tm.Templates[template].Execute(w, data)

	if err != nil {
		log.Println(err)
	}
}

// NewTemplateManager returns a new Template manager.
func NewTemplateManager() Template {
	templates := make(map[string]*template.Template)

	templates["home"] = parseTemplate("home/index")
	templates["about"] = parseTemplate("home/about")
	templates["authenticate"] = parseTemplate("auth/authenticate")

	templateManager := Template{
		Templates: templates,
	}

	return templateManager
}

// parseTemplate compiles a view and panics if the view contains an error.
func parseTemplate(name string) *template.Template {
	name = "app/templates/" + name + ".gohtml"
	layoutPath := "app/templates/layouts/app.gohtml"

	return template.Must(template.ParseFiles(layoutPath, name))
}
