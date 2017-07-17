package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/viewModels"
	"log"
	"net/http"
)

type Home struct {
	Base
}

func (hc *Home) Home(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewHome()
	err := hc.ViewManager.Home.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func (hc *Home) About(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewAbout()
	err := hc.ViewManager.About.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func NewHomeController(viewManager *managers.View) Home {
	return Home{
		Base: NewBaseController(viewManager),
	}
}
