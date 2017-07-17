package viewModels

type About struct {
	Base
}

func NewAbout() About {
	return About{
		Base: NewBase(),
	}
}
