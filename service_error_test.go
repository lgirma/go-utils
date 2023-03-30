package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBusinessError(t *testing.T) {
	err := NewBusinessError("AUTH_FAILED")
	assert.True(t, IsBusinessError(err))
	assert.False(t, IsBusinessError(errors.New("some_error")))
}

func TestBusinessErrorToString(t *testing.T) {
	err := NewBusinessError("AUTH_FAILED")
	assert.Equal(t, "AUTH_FAILED", err.Error())

	err = NewBusinessError("error_code", "details")
	assert.Equal(t, "error_code: details", err.Error())
}

func TestBusinessErrorFromErrNo(t *testing.T) {
	err := NewBusinessErrorFromErrNo(101)
	assert.Equal(t, "101", err.Error())

	err = NewBusinessErrorFromErrNo(101, "details")
	assert.Equal(t, "101: details", err.Error())
}