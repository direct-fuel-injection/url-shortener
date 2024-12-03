package domain

import "errors"

var (
	ErrRequired = errors.New("required value")
	ErrInvalid  = errors.New("invalid value")
	ErrNotFound = errors.New("not found")
	ErrNil      = errors.New("nil data")
)