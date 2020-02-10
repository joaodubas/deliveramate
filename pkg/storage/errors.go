package storage

import "errors"

var (
	ErrorDuplicateID       = errors.New("duplicated id")
	ErrorDuplicateDocument = errors.New("duplicated document")
	ErrorNotFound          = errors.New("id not found")
	ErrorWrongCoverageArea = errors.New("wrong type for coverage area")
	ErrorWrongAddress      = errors.New("wrong type for address")
)
