package bitmultihooks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxError_StringRepresentation(t *testing.T) {
	test := assert.New(t)

	err := syntaxError{1, "a"}
	test.Implements((*error)(nil), err)
	test.Equal("line 1: a", err.Error())
}
