package errors

import "errors"

var (
	ErrInvalidInput = errors.New("request contains invalid input")
)
