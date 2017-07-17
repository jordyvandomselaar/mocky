package viewModels

import "github.com/jordyvandomselaar/mock-backend/app/managers"

type Base struct {
	Urls *managers.Url
}

func NewBase() Base {
	return Base{
		Urls: managers.NewUrlManager(),
	}
}
