package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")

	ErrInvalidCurrency = errors.New("invalid currency")
	ErrAPIError        = errors.New("API error")
	ErrConversionError = errors.New("conversion error")
)

type AppError struct {
	Err     error
	Message string
	Details interface{}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v: %v", e.Message, e.Err, e.Details)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewInvalidCurrencyError(currency string) error {
	return &AppError{
		Err:     ErrInvalidCurrency,
		Message: fmt.Sprintf("Invalid currency: %s", currency),
		Details: currency,
	}
}

func NewAPIError(err error) error {
	return &AppError{
		Err:     ErrAPIError,
		Message: "API error occurred",
		Details: err,
	}
}

func NewConversionError(err error) error {
	return &AppError{
		Err:     ErrConversionError,
		Message: "Conversion error occurred",
		Details: err,
	}
}
