package service

type InvalidPageError struct {
	message string
}

func NewInvalidPageError() error {
	return &InvalidPageError{"page value is invalid"}
}

func (e *InvalidPageError) Error() string {
	return e.message
}

type InvalidMaxResultsError struct {
	message string
}

func NewInvalidMaxResultsError() error {
	return &InvalidMaxResultsError{"maxResults value is invalid"}
}

func (e *InvalidMaxResultsError) Error() string {
	return e.message
}
