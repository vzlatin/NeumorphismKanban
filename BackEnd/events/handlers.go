package events

const (
	NewBoard = 0
	NewColumn = 1
	NewTask = 2
)

type EventHandler func(event Event) error
