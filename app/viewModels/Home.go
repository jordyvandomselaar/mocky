package viewModels

// Home is the homepage viewmodel.
type Home struct {
	Base
}

// NewHome returns a new Home viewmodel.
func NewHome() Home {
	return Home{
		Base: NewBase(),
	}
}
