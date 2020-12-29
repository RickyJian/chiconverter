package utils

import (
	"errors"
)

var (
	// ErrConvertUnimplemented describes convert unimplemented
	ErrConvertUnimplemented = errors.New("convert unimplemented")
)

// Converter is the interface that wraps basic Convert method
type Converter interface {
	Convert() (string, error)
}

// Convert convert chinese
func Convert(val interface{}) (string, error) {
	c, ok := val.(Converter)
	if !ok {
		return "", ErrConvertUnimplemented
	}
	return c.Convert()
}
