package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {

	err := UserError.New("an_error")
	errWithContext := err.AddSingleContext("a_field", "the field is empty")

	expectedContext := map[string]string{"a_field": "the field is empty"}

	assert.Equal(t, UserError, errWithContext.Type())
	assert.Equal(t, expectedContext, errWithContext.Context())
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := New("a custom error")

	errWithContext := err.AddSingleContext("a_field", "the field is empty")

	expectedContext := map[string]string{"a_field": "the field is empty"}

	assert.Equal(t, NoType, errWithContext.Type())
	assert.Equal(t, expectedContext, errWithContext.Context())
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := New("an_error")
	wrappedError := UserError.Wrapf(err, "error %s", "1")

	assert.Equal(t, UserError, wrappedError.Type())
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfNoType(t *testing.T) {
	err := Newf("an_error %s", "2")
	wrappedError := Wrapf(err, "error %s", "1")

	assert.Equal(t, NoType, wrappedError.Type())
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
