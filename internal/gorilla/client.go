package gorilla

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func (c *Client) Send(msg interface{}) error {
	return c.conn.WriteJSON(msg)
}

func (c *Client) Receive() ([]byte, error) {
	_, msg, err := c.conn.ReadMessage()
	return msg, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}
