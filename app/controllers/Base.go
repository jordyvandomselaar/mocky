package controllers

import "github.com/jordyvandomselaar/mock-backend/app/managers"

// Base controller.
type Base struct {
	ViewManager *managers.View
}

// NewBaseController returns a new Base controller.
func NewBaseController(viewManager *managers.View) Base {
	return Base{
		ViewManager: viewManager,
	}
}
