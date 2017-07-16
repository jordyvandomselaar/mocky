package controllers

import (
	"net/http"
	"github.com/jordyvandomselaar/mock-backend/app/helpers"
)

type AuthController struct {
	username string
	password string
	email    string
}

func (ac *AuthController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	viewManager := helpers.NewViewManager()
	viewManager.Authenticate.Execute(w, nil)
}
