package managers

// Url manager allows aliasing of routes and prevents mistakes.
type Url struct {
	Home         string
	Authenticate string
	About        string
}

// NewUrlManager returns a new Url manager.
func NewUrlManager() Url {
	return Url{
		Home:         "/",
		Authenticate: "/authenticate#tab=login",
		About:        "/about",
	}
}
