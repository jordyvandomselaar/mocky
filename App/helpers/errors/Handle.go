package errors

func Handle(error error) {
	if error != nil {
		panic(error)
	}
}
