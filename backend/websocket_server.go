package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	clientsMu sync.Mutex
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *WebSocketServer) Start(addr string) {
	http.HandleFunc("/ws", s.HandleConnections)
	log.Println("WebSocket server started on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("WebSocket server error:", err)
	}
}

func (s *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	// Register new client
	s.clientsMu.Lock()
	s.clients[ws] = true
	s.clientsMu.Unlock()

	log.Println("New WebSocket client connected")

	// Handle messages
	for {
		var msg interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
	}
}

func (s *WebSocketServer) Broadcast(message interface{}) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	for client := range s.clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Println("WebSocket write error:", err)
			client.Close()
			delete(s.clients, client)
		}
	}
}
