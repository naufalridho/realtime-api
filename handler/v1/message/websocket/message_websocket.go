package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"

	"github.com/naufalridho/realtime-api/model"
)

type MessageWebSocketHandlerV1 struct {
	upgrader websocket.Upgrader
	clients map[*websocket.Conn]bool
	broadcast chan model.Message
}

func New() *MessageWebSocketHandlerV1 {
	return &MessageWebSocketHandlerV1{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan model.Message),
	}
}

func (m *MessageWebSocketHandlerV1) ReadMessageWebSocket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "cannot open websocket connection", http.StatusBadRequest)
	}

	defer conn.Close()
	m.clients[conn] = true

	for {
		var message model.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error: %v", err)
			delete(m.clients, conn)
			break
		}
		m.broadcast <- message
	}
}

func (m *MessageWebSocketHandlerV1) WriteMessageWebSocket() {
	for {
		msg := <-m.broadcast
		for client := range m.clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(m.clients, client)
			}
		}
	}
}