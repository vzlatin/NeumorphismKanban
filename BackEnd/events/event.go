package events

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Event struct {
	Type    int             `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type NewColumnEvent struct {
	Type    int       `json:"type"`
	BoardID uuid.UUID `json:"id"`
	Title   string    `json:"title"`
}

// func Deserialize(event []byte) (interface{}, error) {

// 	var payload interface{}
// 	switch data.Type {
// 	case 0:
// 		var board NewBoardEvent
// 		if err := json.Unmarshal([]byte(data.Payload.(string)), &board); err != nil {
// 			return nil, fmt.Errorf("error deserializing into a new board event: %s", err)
// 		}
// 		payload = board
// 	case 1:
// 		var column NewColumnEvent
// 		if err := json.Unmarshal([]byte(data.Payload.(string)), &column); err != nil {
// 			return nil, fmt.Errorf("error deserializing into a new column event: %s", err)
// 		}
// 		payload = column
// 	// Add cases for other event types as needed
// 	default:
// 		return nil, fmt.Errorf("unsupported event type: %d", data.Type)
// 	}

// 	return payload, nil
// }
