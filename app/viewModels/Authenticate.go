package viewModels

// Authenticate is the authenticate page viewmodel.
type Authenticate struct {
	Base
}

// NewAuthenticate returns a new Authenticate viewmodel.
func NewAuthenticate() Authenticate {
	return Authenticate{
		Base: NewBase(),
	}
}
