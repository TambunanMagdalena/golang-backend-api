package customerror

import (
	"errors"
	"net/http"
)

type TrackableError struct {
	err error
}

func (e TrackableError) Error() string {
	return e.err.Error()
}

func (e TrackableError) Unwrap() error {
	return e.err
}

func New(msg string) error {
	return errors.New(msg)
}

func GetStatusCode(err error) int {
	var nf NotFoundError
	var br BadRequestError
	var ie InternalServiceError

	if errors.As(err, &nf) {
		return http.StatusNotFound
	}
	if errors.As(err, &br) {
		return http.StatusBadRequest
	}
	if errors.As(err, &ie) {
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}