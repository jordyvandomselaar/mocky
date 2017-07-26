package controllers

import "github.com/jordyvandomselaar/mock-backend/app/managers"

// Base controller.
type Base struct {
	TemplateManager managers.Template
}

// NewBaseController returns a new Base controller.
func NewBaseController(templateManager managers.Template) Base {
	return Base{
		TemplateManager: templateManager,
	}
}
