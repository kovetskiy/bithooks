package bithooks

import "strings"

func Encode(hooks Hooks) string {
	encoded := []string{}

	for _, hook := range hooks {
		encoded = append(encoded, hook.Name+"@"+hook.ID)
		for _, arg := range hook.Args {
			encoded = append(encoded, " "+arg)
		}

		// empty line between hooks
		encoded = append(encoded, "")
	}

	return strings.Join(encoded, "\n")
}
