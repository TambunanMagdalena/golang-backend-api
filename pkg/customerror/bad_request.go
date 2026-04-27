package customerror

import "fmt"

type badRequest struct {
	TrackableError
}

type BadRequestError interface {
	error
	IsBadRequestError() bool
}

func (e *badRequest) IsBadRequestError() bool { return true }

func NewBadRequestErrorf(format string, data ...interface{}) error {
	return &badRequest{
		TrackableError{fmt.Errorf(format, data...)},
	}
}

func NewBadRequestError(message string) error {
	return &badRequest{
		TrackableError{fmt.Errorf(message)},
	}
}