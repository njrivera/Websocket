package gorilla

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/Websocket/pkg/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	clientChan chan types.Client
	clients    map[string]*Client
	sync.RWMutex
}

func NewServer(url, port string) (*Server, error) {
	server := &Server{
		clients:    map[string]*Client{},
		clientChan: make(chan types.Client),
	}

	http.HandleFunc(url, server.handler)
	go http.ListenAndServe(":"+port, nil)

	return server, nil
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr

	s.Lock()
	if _, ok := s.clients[addr]; ok {
		log.Printf("Connection already established for address: %s", addr)
		w.WriteHeader(http.StatusBadRequest)
		s.Unlock()
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		s.Unlock()
		return
	}

	c := &Client{conn: conn}
	s.clients[addr] = c
	s.Unlock()
	s.clientChan <- c
}

// Broadcast sends a given message to all connected clients
func (s *Server) Broadcast(msg interface{}) {
	// Put clients into slice first to lock the clients map for as little time as possible
	s.RLock()
	clients := []*Client{}
	for _, c := range s.clients {
		clients = append(clients, c)
	}
	s.RUnlock()

	for _, c := range clients {
		if err := c.Send(msg); err != nil {
			log.Printf("Unable to send message to client %s, Error: %s", c.conn.RemoteAddr().String(), err.Error())
		}
	}
}

// ListenForNewClients returns a channel that all new ws clients are pushed into
func (s *Server) ListenForNewClients() <-chan types.Client {
	return s.clientChan
}
