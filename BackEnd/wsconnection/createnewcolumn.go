package wsconnection

import (
	// "context"
	// "os"

	// "github.com/google/uuid"
	// "github.com/joho/godotenv"
	// "github.com/vzlatin/NeumorphismKanban/internal/database"
	"fmt"

	"github.com/vzlatin/NeumorphismKanban/events"
)

func CreateNewColumn(event events.Event, c *Client) error {
	fmt.Println("Printing from new column : ", event.Payload)
	// apiConfig.DB.CreateBoard(ctx, database.CreateBoardParams{
	// 	ID:    uuid.New(),
	// 	Title: "Poop",
	// })
	return nil
}
