package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/viewModels"
	"net/http"
)

// Home controller handles the main pages.
type Home struct {
	Base
}

// Home renders the homepage.
func (hc Home) Home(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewHome()
	hc.TemplateManager.Render("home", w, data)
}

// About renders the about page.
func (hc Home) About(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewAbout()
	hc.TemplateManager.Render("about", w, data)
}

// NewHomeController returns a new Home controller.
func NewHomeController(viewManager managers.Template) Home {
	return Home{
		Base: NewBaseController(viewManager),
	}
}
