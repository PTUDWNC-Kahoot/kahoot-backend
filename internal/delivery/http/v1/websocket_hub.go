package v1

import (
	"encoding/json"
	"log"
)

type Boardcast struct {
	roomID uint32
	data   []byte
}

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Boardcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister    chan *Client
	presentations map[uint32]map[string]interface{}
}

func newHub() *Hub {
	return &Hub{
		broadcast:     make(chan Boardcast),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		presentations: make(map[uint32]map[string]interface{}),
	}
}

type Message struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case b := <-h.broadcast:
			m := Message{}
			err := json.Unmarshal(b.data, &m)
			if err != nil {
				log.Println("failed to unmarshal data", err)
				continue
			}

			if h.presentations[b.roomID] == nil {
				h.presentations[b.roomID] = map[string]interface{}{}
			}

			h.presentations[b.roomID][m.Action] = m.Payload // set current state of presentation

			for client := range h.clients {
				if client.roomID == b.roomID {
					select {
					case client.send <- b.data:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}

			}
		}
	}
}
