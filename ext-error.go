package exterrors

import (
	"encoding/json"
	"net/http"
)

// ExtError describes extended error.
// Can be used in the same way as basic error.
type ExtError interface {
	error
	json.Marshaler

	// HTTPCode contains info about HTTP code
	// that could be sent in the status of the HTTP response.
	HTTPCode() int

	// ErrMessage contains error message if any.
	ErrMessage() string

	// ErrField describes in which field an error has occurred if any.
	// For example JSON field in the HTTP request body.
	// Typically is used for pointing an invalid field during validation.
	ErrField() string
}

// Error contains extended error information.
type Error struct {

	// Code contains info about HTTP code
	// that could be sent in the status of the HTTP response.
	Code int `json:"-"`

	// Message contains error message if any.
	Message string `json:"message,omitempty"`

	// Description contains detailed error information if any.
	Description string `json:"description,omitempty"`

	// Field describes in which field an error has occurred if any.
	// For example JSON field in the HTTP request body.
	// Typically is used for pointing an invalid field during validation.
	Field string `json:"filed,omitempty"`
}

// Error unifying Error with Go's error interface.
func (e Error) Error() string {
	return e.Description
}

func (e Error) HTTPCode() int {
	return e.Code
}

func (e Error) ErrMessage() string {
	return e.Message
}

func (e Error) ErrField() string {
	return e.Field
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}

// NewError returns new ExtError with filled Message, Description and Field.
func NewError(err error, code int, message, field string) ExtError {
	if err == nil {
		return nil
	}

	return &Error{
		Code:        code,
		Message:     message,
		Description: err.Error(),
		Field:       field,
	}
}

func NewBadRequestError(err error) ExtError {
	return NewError(err, http.StatusBadRequest, "Bad Request", "")
}

func NewUnauthorizedError(err error) ExtError {
	return NewError(err, http.StatusUnauthorized, "Unauthorized", "")
}

func NewForbiddenError(err error) ExtError {
	return NewError(err, http.StatusForbidden, "Forbidden", "")
}

func NewNotFoundError(err error, field string) ExtError {
	return NewError(err, http.StatusNotFound, "Not Found", field)
}

func NewUnprocessableEntityError(err error, field string) ExtError {
	return NewError(err, http.StatusUnprocessableEntity, "Unprocessable Entity", field)
}

func NewInternalServerErrorError(err error) ExtError {
	return NewError(err, http.StatusInternalServerError, "Internal Server Error", "")
}
