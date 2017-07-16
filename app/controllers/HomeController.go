package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/helpers"
	"net/http"
)

type HomeController struct{}

func (hc *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	viewManager := helpers.NewViewManager()
	viewManager.Home.Execute(w, nil)
}
