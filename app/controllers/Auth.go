package controllers

import (
	"github.com/jordyvandomselaar/mock-backend/app/managers"
	"github.com/jordyvandomselaar/mock-backend/app/viewModels"
	"net/http"
)

// Auth handles the authentication of users.
type Auth struct {
	Base

	Username string
	Password string
	Email    string
}

// ShowLogin renders the authentication page.
func (ac Auth) ShowAuthenticate(w http.ResponseWriter, r *http.Request) {
	data := viewModels.NewAuthenticate()
	ac.TemplateManager.Render("authenticate", w, data)
}

// NewAuthController returns a new instance of the Auth controller.
func NewAuthController(viewManager managers.Template) Auth {
	return Auth{
		Base: NewBaseController(viewManager),
	}
}
