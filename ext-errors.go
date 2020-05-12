package exterrors

import (
	"encoding/json"
	"strings"
)

type ExtErrors interface {
	error
	json.Marshaler

	// Add adds new ExtError to the slice.
	Add(...ExtError)

	// Errors returns the slice of ExtError for future processing.
	Errors() []ExtError

	// Len returns current number of added errors.
	Len() int
}

// Errors represents multiple errors occurred.
// Useful for validation errors when required to show all errors at once.
type Errors struct {
	Errs []ExtError `json:"errors"`
}

// NewExtErrors returns usage ready instance of ExtErrors.
func NewExtErrors() ExtErrors {
	return NewExtErrorsWithCap(1)
}

// NewExtErrorsWithCap returns instance of ExtErrors
// with specified capacity of underlying slice.
func NewExtErrorsWithCap(capacity int) ExtErrors {
	return &Errors{Errs: make([]ExtError, 0, capacity)}
}

func (errs *Errors) Add(err ...ExtError) {
	errs.Errs = append(errs.Errs, err...)
}

func (errs Errors) Len() int {
	return len(errs.Errs)
}

func (errs Errors) MarshalJSON() ([]byte, error) {
	e := make([]Error, errs.Len())
	for i, err := range errs.Errs {
		e[i] = Error{
			Message:     err.ErrMessage(),
			Description: err.ErrDescription(),
			Field:       err.ErrField(),
		}
	}

	return json.Marshal(e)
}

func (errs Errors) Error() string {
	var errors = make([]string, len(errs.Errs))

	for i := range errs.Errs {
		errors[i] = errs.Errs[i].Error()
	}

	return strings.Join(errors, "; ")
}

func (errs Errors) Errors() []ExtError {
	return errs.Errs
}
