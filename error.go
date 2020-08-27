package go_utils

// PanicWhenHasError panics when the error is not nil
func PanicWhenHasError(err error) {
	if err != nil {
		panic(err)
	}
}
