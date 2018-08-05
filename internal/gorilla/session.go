package gorilla

import "github.com/gorilla/websocket"

type session struct {
	Conn *websocket.Conn
}

func (sess *session) receive() ([]byte, error) {
	_, msg, err := sess.Conn.ReadMessage()
	return msg, err
}

func (sess *session) send(msg interface{}) error {
	return sess.Conn.WriteJSON(msg)
}

func (sess *session) close() error {
	return sess.Conn.Close()
}
