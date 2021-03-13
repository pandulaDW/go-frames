package base

// ApplyFunc takes in any value and return another value alongside an error if encountered
type ApplyFunc func(val interface{}) (interface{}, error)
