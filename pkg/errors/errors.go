package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidCustomerID = errors.New("invalid customer id")
	ErrCustomerNotFound  = errors.New("customer not found")
	ErrGeneral           = errors.New("there was an error occurred")
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}
