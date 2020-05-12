package exterrors

import (
	"encoding/json"
	"testing"
)

func TestError_ErrDescription(t *testing.T) {
	const want = "test description"

	err := Error{
		Description: want,
	}

	if got := err.ErrDescription(); got != want {
		t.Errorf("ErrDescription() = %v, want %v", got, want)
	}
}

func TestError_ErrField(t *testing.T) {
	const want = "test field"

	err := Error{
		Field: want,
	}

	if got := err.ErrField(); got != want {
		t.Errorf("ErrField() = %v, want %v", got, want)
	}
}

func TestError_ErrMessage(t *testing.T) {
	const want = "test message"

	err := Error{
		Message: want,
	}

	if got := err.ErrMessage(); got != want {
		t.Errorf("ErrMessage() = %v, want %v", got, want)
	}
}

func TestError_Error(t *testing.T) {
	const want = "test message: test description"

	err := Error{
		Message:     "test message",
		Description: "test description",
	}

	if got := err.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestError_HTTPCode(t *testing.T) {
	const want = 500

	err := Error{
		Code: want,
	}

	if got := err.HTTPCode(); got != want {
		t.Errorf("HTTPCode() = %v, want %v", got, want)
	}
}

func TestError_MarshalJSON(t *testing.T) {
	err := Error{
		Code:        500,
		Message:     "test message",
		Description: "test description",
		Field:       "test field",
	}

	got, e := json.Marshal(err)
	if e != nil {
		t.Errorf("json.Marshal(), err: %s", err.Error())
	}

	var err2 Error
	e = json.Unmarshal(got, &err2)
	if e != nil {
		t.Errorf("json.Unmarshal(), err: %s", err.Error())
	}

	if err.Message != err2.Message {
		t.Errorf("invalid message, got: %v, want: %v", err2.Message, err.Message)
	}

	if err.Field != err2.Field {
		t.Errorf("invalid field, got: %v, want: %v", err2.Field, err.Field)
	}

	if err.Description != err2.Description {
		t.Errorf("invalid description, got: %v, want: %v", err2.Description, err.Description)
	}
}
