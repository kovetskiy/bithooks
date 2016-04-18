package bithooks

import "fmt"

type ErrorHookExists struct {
	Name string
	ID   string
}

func (err ErrorHookExists) Error() string {
	return fmt.Sprintf("hook %s@%s already exists", err.Name, err.ID)
}

type Hooks []*Hook

type Hook struct {
	Name string
	ID   string
	Data string
}

func (hooks *Hooks) Append(hook *Hook) error {
	_, ok := hooks.Get(hook.Name, hook.ID)
	if ok {
		return ErrorHookExists{Name: hook.Name, ID: hook.ID}
	}

	*hooks = append(*hooks, hook)

	return nil
}

func (hooks *Hooks) Get(name, id string) (*Hook, bool) {
	for _, hook := range *hooks {
		if hook.Name == name && hook.ID == id {
			return hook, true
		}
	}

	return nil, false
}

func (hooks *Hooks) Delete(name, id string) {
	for i, hook := range *hooks {
		if hook.Name == name && hook.ID == id {
			*hooks = append((*hooks)[:i], (*hooks)[i+1:]...)
		}
	}
}
