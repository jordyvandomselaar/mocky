package controllers

import "github.com/jordyvandomselaar/mock-backend/app/managers"

type Base struct {
	ViewManager *managers.View
}

func NewBaseController(viewManager *managers.View) Base {
	return Base{
		ViewManager: viewManager,
	}
}
