package events

import "encoding/json"

const (
	NewBoard  = 0
	NewColumn = 1
	NewTask   = 2
)

type Event struct {
	Type    int             `json:"type"`
	Payload json.RawMessage `json:"payload"`
}