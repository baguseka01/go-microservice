package business

import "errors"

var (
	// ErrInternalServeer Error caused by system error
	ErrInternalServeer = errors.New("internal server error")

	// ErrHasBeenModified Error when update item that has been modified
	ErrHasBeenModified = errors.New("data has been modified")

	// ErrNotFound Error when item is not found
	ErrNotFound = errors.New("data was not found")

	// ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("given spec is not valid")
)
