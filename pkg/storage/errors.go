package storage

import "errors"

var (
	ErrorDuplicateID       = errors.New("duplicated id")
	ErrorDuplicateDocument = errors.New("duplicated document")
	ErrorNotFound          = errors.New("id not found")
	ErrorWrongAddress      = errors.New("wrong type for address")
)
