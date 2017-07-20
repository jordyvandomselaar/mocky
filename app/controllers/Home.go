package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/viewModels"
	"log"
	"net/http"
)

// Home controller handles the main pages.
type Home struct {
	Base
}

// Home renders the homepage.
func (hc *Home) Home(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewHome()
	err := hc.ViewManager.Home.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

// About renders the about page.
func (hc *Home) About(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewAbout()
	err := hc.ViewManager.About.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

// NewHomeController returns a new Home controller.
func NewHomeController(viewManager *managers.View) Home {
	return Home{
		Base: NewBaseController(viewManager),
	}
}
