package customerror

import "fmt"

type internalService struct {
	TrackableError
}

type InternalServiceError interface {
	error
	IsInternalError() bool
}

func (e *internalService) IsInternalError() bool { return true }

func NewInternalServiceErrorf(format string, data ...interface{}) error {
	return &internalService{
		TrackableError{fmt.Errorf(format, data...)},
	}
}

func NewInternalServiceError(message string) error {
	return &internalService{
		TrackableError{fmt.Errorf(message)},
	}
}