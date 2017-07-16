package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/App/helpers"
	"net/http"
)

type HomeController struct{}

func (hc *HomeController) Home(writer http.ResponseWriter, request *http.Request) {
	view := helpers.View{Path: "App/views/home/index"}
	view.Render(writer, nil)
}
