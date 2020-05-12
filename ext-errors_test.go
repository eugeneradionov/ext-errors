package exterrors

import (
	"encoding/json"
	"testing"
)

func TestErrors_MarshalJSON(t *testing.T) {
	err := Error{
		Code:        500,
		Message:     "test message",
		Description: "test description",
		Field:       "test field",
	}

	errs := Errors{
		Errs: make([]ExtError, 0),
	}
	errs.Add(err)

	got, e := json.Marshal(errs)
	if e != nil {
		t.Errorf("json.Marshal(), err: %s", err.Error())
	}

	var jsonErrs []Error
	e = json.Unmarshal(got, &jsonErrs)
	if e != nil {
		t.Errorf("json.Unmarshal(), err: %s", e.Error())
	}

	newErrs := make([]ExtError, len(jsonErrs))
	for i := range jsonErrs {
		newErrs[i] = jsonErrs[i]
	}

	errs2 := Errors{
		Errs: newErrs,
	}

	err2 := errs2.Errs[0]
	if err.Message != err2.ErrMessage() {
		t.Errorf("invalid message, got: %v, want: %v", err2.ErrMessage(), err.Message)
	}

	if err.Field != err2.ErrField() {
		t.Errorf("invalid field, got: %v, want: %v", err2.ErrField(), err.Field)
	}

	if err.Description != err2.ErrDescription() {
		t.Errorf("invalid description, got: %v, want: %v", err2.ErrDescription(), err.Description)
	}
}
