package bitmultihooks

type Hooks []Hook

type Hook struct {
	Name string
	ID   string
	Data string
}
