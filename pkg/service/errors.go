package service

type InvalidPageError struct {
	Err error
}

func (e *InvalidPageError) Error() string {
	return e.Err.Error()
}