package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/helpers"
	"net/http"
	"log"
)

type AuthController struct {
	username string
	password string
	email    string
}

func (ac *AuthController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	viewManager := helpers.NewViewManager()
	err := viewManager.Authenticate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
