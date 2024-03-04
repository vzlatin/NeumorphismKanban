package wsconnection

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/vzlatin/NeumorphismKanban/events"
	"github.com/vzlatin/NeumorphismKanban/dbhandlers"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		ReadBufferSize:   2048,
		WriteBufferSize:  2048,
	}
)

type Manager struct {
	clients  ClientList
	handlers map[int]events.EventHandler
	sync.RWMutex
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[int]events.EventHandler),
	}
	m.setupEventHandlers()
	return m
}

// This function is called whenever a request from the client is received
func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	websocketUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Couldn't establish a websocket connection: %s", err)
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[c] = true
}

func (m *Manager) removeClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[c]; ok {
		c.connection.Close()
		delete(m.clients, c)
	}
}

func (m *Manager) setupEventHandlers() {
	m.handlers[events.NewBoard] = dbhandlers.CreateNewBoard
	m.handlers[events.NewColumn] = dbhandlers.CreateNewColumn
}

func (m *Manager) route(event events.Event) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}
