package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/helpers"
	"net/http"
	"log"
)

type HomeController struct{}

func (hc *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	viewManager := helpers.NewViewManager()
	err := viewManager.Home.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (hc *HomeController) About(w http.ResponseWriter, r *http.Request) {
	viewManager := helpers.NewViewManager()
	err := viewManager.About.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}