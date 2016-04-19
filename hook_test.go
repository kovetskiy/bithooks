package bithooks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHooks_Get_ReturnsHookAndTrueIfFound(t *testing.T) {
	test := assert.New(t)

	hooks := Hooks{
		{"x", "y", nil},
	}
	actualHook, ok := hooks.Get("x", "y")

	test.True(ok)
	test.Equal(hooks[0], actualHook)
}

func TestHooks_Append_AppendsHook(t *testing.T) {
	test := assert.New(t)

	expectedHook := &Hook{"x", "y", nil}

	hooks := Hooks{}
	hooks.Append(expectedHook)

	if test.Len(hooks, 1) {
		test.Equal(expectedHook, hooks[0])
	}
}

func TestHook_Delete(t *testing.T) {
	test := assert.New(t)

	hooks := Hooks{
		{"x", "y", nil},
	}
	hooks.Delete("x", "y")

	test.Len(hooks, 0)
}
