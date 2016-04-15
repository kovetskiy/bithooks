package bitmultihooks

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
	errSyntaxUnexpectedHookData = "unexpected hook data"
	errSyntaxDefine             = "hook should be " +
		"defined as <name>@<identifier>"
	errSyntaxRedefine = "cannot redefine hook, " +
		"hook with same <name>@<identifier> already defined"
)

func Decode(data string) (Hooks, error) {
	lines := strings.Split(data, "\n")

	var (
		hooks   = Hooks{}
		hook    = Hook{}
		defines = map[string]struct{}{}
	)

	for index, line := range lines {
		switch strings.HasPrefix(line, " ") {
		case true:
			if hook.Name == "" {
				return hooks,
					syntaxError{index + 1, errSyntaxUnexpectedHookData}
			}

			line = strings.TrimPrefix(line, " ")
			if hook.Data == "" {
				hook.Data = line
			} else {
				hook.Data += "\n" + line
			}

		case false:
			if hook.Name != "" {
				defines[hook.Name+"@"+hook.ID] = struct{}{}
				hooks = append(hooks, hook)
				hook = Hook{}
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

			_, redefine := defines[hook.Name+"@"+hook.ID]
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
