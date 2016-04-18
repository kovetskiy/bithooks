package bithooks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHooks_Get_ReturnsHookAndTrueIfFound(t *testing.T) {
	test := assert.New(t)

	hooks := Hooks{
		{"x", "y", ""},
	}
	actualHook, ok := hooks.Get("x", "y")

	test.True(ok)
	test.Equal(hooks[0], actualHook)
}

func TestHooks_Append_AppendsHook(t *testing.T) {
	test := assert.New(t)

	expectedHook := &Hook{"x", "y", ""}

	hooks := Hooks{}
	hooks.Append(expectedHook)

	if test.Len(hooks, 1) {
		test.Equal(expectedHook, hooks[0])
	}
}

func TestHook_Delete(t *testing.T) {
	test := assert.New(t)

	hooks := Hooks{
		{"x", "y", ""},
	}
	hooks.Delete("x", "y")

	test.Len(hooks, 0)
}
