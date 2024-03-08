package service

type InvalidPageError struct {
	message string
}

func NewInvalidPageError() error {
	return InvalidPageError{"page value is invalid"}
}

func (e InvalidPageError) Error() string {
	return e.message
}

type InvalidMaxResultsError struct {
	message string
}

func (e InvalidMaxResultsError) Error() string {
	return e.message
}

func NewInvalidMaxResultsError() error {
	return InvalidMaxResultsError{"maxResults value is invalid"}
}

type ResourceNotExistError struct {
	message string
}

func NewResourceNotExistError() error {
	return ResourceNotExistError{"Resource not exists"}
}

func (e ResourceNotExistError) Error() string {
	return e.message
}

type ResourceAlreadyExistError struct {
	message string
}

func NewResourceAlreadyExistError() error {
	return ResourceAlreadyExistError{"Resource already exists"}
}

func (e ResourceAlreadyExistError) Error() string {
	return e.message
}
