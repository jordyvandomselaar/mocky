package viewModels

type Home struct {
	Base
}

func NewHome() Home {
	return Home{
		Base: NewBase(),
	}
}
