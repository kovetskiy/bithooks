package bithooks

import (
	"fmt"
	"strings"
)

type syntaxError struct {
	line int
	text string
}

func (err syntaxError) Error() string {
	return fmt.Sprintf("line %d: %s", err.line, err.text)
}

var (
	errSyntaxUnexpectedHookArgs = "unexpected hook args"
	errSyntaxDefine             = "hook should be " +
		"defined as <name>@<identifier>"
	errSyntaxRedefine = "cannot redefine hook, " +
		"hook with same <name>@<identifier> already defined"
)

func Decode(args string) (Hooks, error) {
	lines := strings.Split(args, "\n")

	var (
		hooks = Hooks{}
		hook  = &Hook{}
	)

	for index, line := range lines {
		switch strings.HasPrefix(line, " ") {
		case true:
			if hook.Name == "" {
				return hooks,
					syntaxError{index + 1, errSyntaxUnexpectedHookArgs}
			}

			hook.Args = append(hook.Args, strings.TrimPrefix(line, " "))

		case false:
			if hook.Name != "" {
				hooks.Append(hook)
				hook = &Hook{}
			}

			if line == "" {
				continue
			}

			subject := strings.Split(line, "@")
			if len(subject) != 2 {
				return hooks,
					syntaxError{index + 1, errSyntaxDefine}
			}

			hook.Name = subject[0]
			hook.ID = subject[1]

			_, redefine := hooks.Get(hook.Name, hook.ID)
			if redefine {
				return hooks, syntaxError{index + 1, errSyntaxRedefine}
			}
		}
	}

	if hook.Name != "" {
		hooks = append(hooks, hook)
	}

	return hooks, nil
}
