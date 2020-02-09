package storage

import "errors"

var (
	ErrorDuplicateID       = errors.New("duplicated id")
	ErrorDuplicateDocument = errors.New("duplicated document")
	ErrorWrongAddress      = errors.New("wrong type for address")
)
