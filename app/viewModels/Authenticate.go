package viewModels

type Authenticate struct {
	Base
}

func NewAuthenticate() Authenticate {
	return Authenticate{
		Base: NewBase(),
	}
}
