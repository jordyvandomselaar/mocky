package managers

type Url struct {
	Home         string
	Authenticate string
	About        string
}

func NewUrlManager() *Url {
	return &Url{
		Home:         "/",
		Authenticate: "/authenticate#tab=login",
		About:        "/about",
	}
}
