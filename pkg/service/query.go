package service

type Query struct {
	Filters   map[string]interface{}
	Page      int
	MaxResults int
}