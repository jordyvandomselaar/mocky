package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/viewModels"
	"log"
	"net/http"
)

type Auth struct {
	Base

	Username string
	Password string
	Email    string
}

func (ac *Auth) ShowLogin(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewAuthenticate()
	err := ac.ViewManager.Authenticate.Execute(w, data)

	if err != nil {
		log.Println(err)
	}
}

func NewAuthController(viewManager *managers.View) Auth {
	return Auth{
		Base: NewBaseController(viewManager),
	}
}
