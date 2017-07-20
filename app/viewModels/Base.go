package viewModels

import "github.com/jordyvandomselaar/mock-backend/app/managers"

// Base is the base viewmodel.
type Base struct {
	Urls *managers.Url
}

// NewBase returns a new Base viewmodel.
func NewBase() Base {
	return Base{
		Urls: managers.NewUrlManager(),
	}
}
