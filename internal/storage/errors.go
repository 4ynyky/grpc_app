package storage

import "errors"

var (
	ErrConnNotEstablished = errors.New("connection wasn't established")

	ErrNotFound = errors.New("not found")

	ErrItemInvalid = errors.New("Item is invalid")
)
