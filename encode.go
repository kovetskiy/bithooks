package bithooks

import "strings"

func Encode(hooks Hooks) string {
	encoded := []string{}

	for _, hook := range hooks {
		encoded = append(encoded, hook.Name+"@"+hook.ID)
		if len(hook.Args) > 0 {
			lines := strings.Split(hook.Args, "\n")
			for _, line := range lines {
				encoded = append(encoded, " "+line)
			}
		}

		encoded = append(encoded, "")
	}

	return strings.Join(encoded, "\n")
}
