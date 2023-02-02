package storage

import "errors"

var (
	ErrConnNotEstablished = errors.New("connection wasn't established")

	ErrNotFound = errors.New("not found")
)
