package viewModels

// About is the about page viewmodel.
type About struct {
	Base
}

// NewAbout returns a new About view model.
func NewAbout() About {
	return About{
		Base: NewBase(),
	}
}
