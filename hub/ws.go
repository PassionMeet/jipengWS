package hub

import (
	"encoding/json"
	"fmt"

	"github.com/cmfunc/jipengWS/protocol"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.remoteAddr] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.remoteAddr]; ok {
				delete(h.clients, client.remoteAddr)
				close(client.send)
			}
		case message := <-h.broadcast:
			// 这里的逻辑，是将消息分发给了所有的客户端
			// 这里应该是发送给指定的user_id对应的连接，不是遍历
			fmt.Println(string(message))
			cliMsg := protocol.MessageBase{}
			err := json.Unmarshal(message, &cliMsg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if client, ok := h.clients[cliMsg.RmtAddr]; ok {
				client.send <- message
			}
		}
	}
}
