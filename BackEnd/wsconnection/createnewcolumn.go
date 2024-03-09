package wsconnection

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/vzlatin/NeumorphismKanban/events"
	"github.com/vzlatin/NeumorphismKanban/internal/database"
)

type Column struct {
	BoardID uuid.UUID `json:"boardId"`
	Title   string    `json:"title"`
}

func CreateNewColumn(event events.Event, c *Client) error {

	var column Column
	if err := json.Unmarshal([]byte(event.Payload), &column); err != nil {
		return fmt.Errorf("error deserializing into a new column event: %s", err)
	}

	// Create a new column
	apiConfig, ctx := GetApiConfig()
	newcolumn, err := apiConfig.DB.CreateColumn(ctx, database.CreateColumnParams{
		ID:      uuid.New(),
		Boardid: column.BoardID,
		Title:   column.Title,
	})
	if err != nil {
		return fmt.Errorf("error creating a new column: %s", err)
	}

	// // Create a new Board - Column relationship

	// newBoardColumnRelationship, err := apiConfig.DB.CreateBoardIdColumns(ctx, database.CreateBoardIdColumnsParams{
	// 	ID:       uuid.New(),
	// 	Boardid:  column.BoardID,
	// 	Columnid: newcolumn.ID,
	// })
	// if err != nil {
	// 	return fmt.Errorf("error creating a new board - column relationship: %s", err)
	// }
	// // Just create the relationship, no need to return it.
	// _ = newBoardColumnRelationship

	var outgoingMessage events.Event
	outgoingMessagePyload, err := json.Marshal(newcolumn)
	if err != nil {
		return fmt.Errorf("error creating an outgoing column message: %s", err)
	}

	outgoingMessage.Type = 1
	outgoingMessage.Payload = outgoingMessagePyload

	for client := range c.manager.clients {
		client.bottleneck <- outgoingMessage
	}

	return nil
}
