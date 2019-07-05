package errors

type draftError struct {
	errorCode int
	errorMsg string
}
func  ReturnError(err string) draftError{
	return draftError{
		errorCode: 1,
		errorMsg:  "this is frst error",
	}
}
