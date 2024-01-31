package service

type InvalidFilterError struct {
	Err error
}

func (e *InvalidFilterError) Error() string {
	return e.Err.Error()
}

type InvalidPageError struct {
	Err error
}

func (e *InvalidPageError) Error() string {
	return e.Err.Error()
}