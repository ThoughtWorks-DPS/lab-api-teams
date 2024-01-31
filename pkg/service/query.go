package service

type Query struct {
	Filters   map[string]string
	Page      int
	MaxResult int
}