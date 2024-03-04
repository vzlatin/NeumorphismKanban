package wsconnection

import (
	"encoding/json"
	"log"

	"github.com/vzlatin/NeumorphismKanban/events"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	// Used to mitigate spamming, because gorilla allows only one concurrent writer
	bottleneck chan events.Event
}

func NewClient(c *websocket.Conn, m *Manager) *Client {
	return &Client{
		connection: c,
		manager:    m,
		bottleneck: make(chan events.Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		_, event, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Connection closed unexpectedly: %s", err)
			}
			break
		}

		var data events.Event
		if err := json.Unmarshal(event, &data); err != nil {
			log.Printf("Error handling the message: %s", err)
		}

		if err := c.manager.route(data, c); err != nil {
			log.Printf("Error handling the message: %s", err)
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		select {
		case msg, ok := <-c.bottleneck:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("Channel is down, closing connection: %s", err)
				}
				return
			}

			data, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error serializing the message:, %s", err)
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("Error seding the message: %s", err)
			}
			log.Println("Message Sent")
		}
	}
}
