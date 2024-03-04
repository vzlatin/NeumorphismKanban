package wsconnection

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/vzlatin/NeumorphismKanban/events"
	"github.com/vzlatin/NeumorphismKanban/internal/database"
)

type NewBoard struct {
	Title string `json:"title"`
}

func CreateNewBoard(event events.Event, c *Client) error {

	var board NewBoard
	if err := json.Unmarshal([]byte(event.Payload), &board); err != nil {
		return fmt.Errorf("error deserializing into a new board event: %s", err)
	}

	apiConfig, ctx := GetApiConfig()
	apiConfig.DB.CreateBoard(ctx, database.CreateBoardParams{
		ID:    uuid.New(), // Think about this one
		Title: board.Title,
	})

	for client := range c.manager.clients {
		client.bottleneck <- event
	}

	return nil
}
