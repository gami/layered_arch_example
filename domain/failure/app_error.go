package failure

import "github.com/pkg/errors"

type AppError struct {
	Code  errcode
	cause error
}

type errcode string

const (
	ErrNotFound  errcode = "NotFound"
	ErrInvalid   errcode = "InvalidArgument"
	ErrForbidden errcode = "Forbidden"
	ErrConflict  errcode = "Conflict"
)

func NotFound(err error) *AppError {
	return &AppError{
		Code:  ErrNotFound,
		cause: err,
	}
}

func Invalid(err error) *AppError {
	return &AppError{
		Code:  ErrInvalid,
		cause: err,
	}
}

func Forbidden(err error) *AppError {
	return &AppError{
		Code:  ErrForbidden,
		cause: err,
	}
}

func Conflict(err error) *AppError {
	return &AppError{
		Code:  ErrConflict,
		cause: err,
	}
}

func (e *AppError) Error() string {
	return errors.Wrap(e.cause, string(e.Code)).Error()
}

func (e *AppError) Unwrap() error {
	return e.cause
}
