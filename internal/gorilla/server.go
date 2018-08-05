package gorilla

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	//TODO make concurrent
	clientSessions map[int]*session
	//TODO replace with something nicer
	id     int
	idChan chan int
}

func NewWsServer(url string, port string) (*Server, <-chan int) {
	server := &Server{
		clientSessions: map[int]*session{},
		idChan:         make(chan int),
	}
	http.HandleFunc(url, server.handler)
	go http.ListenAndServe(":"+port, nil)

	return server, server.idChan
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	server.clientSessions[server.id] = &session{Conn: conn}
	server.id++
	server.idChan <- server.id
}

func (server *Server) Send(id int, msg interface{}) error {
	if sess, ok := server.clientSessions[id]; ok {
		return sess.send(msg)
	}

	return errors.New("Session does not exist")
}

func (server *Server) Receive(id int) ([]byte, error) {
	if sess, ok := server.clientSessions[id]; ok {
		return sess.receive()
	}

	return nil, errors.New("Session does not exist")
}

func (server *Server) Close(id int) error {
	if sess, ok := server.clientSessions[id]; ok {
		return sess.close()
	}

	return errors.New("Session does not exist")
}
