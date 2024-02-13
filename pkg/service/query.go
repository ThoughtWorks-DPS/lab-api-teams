package service

const MAX_RESULTS = 25

type Query struct {
	Filters   map[string]interface{}
	Page      int
	MaxResults int
}
