package websocket

import (
	"fmt"
	"log"
	"os"

	"github.com/Websocket/internal/gorilla"
	websockettypes "github.com/Websocket/pkg/types"
)

// Websocket types
const (
	Gorilla = "GORILLA"
)

type ErrUnknownWebsocketType struct {
	wsType string
}

func (e ErrUnknownWebsocketType) Error() string {
	return fmt.Sprintf("Unknown websocket type: %s", e.wsType)
}

func NewWsServer(wsType, url, port string) (websockettypes.Server, error) {
	return newServer(wsType, url, port)
}

func NewDefaultWsServer(url, port string) (websockettypes.Server, error) {
	wsType := os.Getenv("WS_TYPE")
	if wsType == "" {
		log.Printf("WS_TYPE not set - defaulting to gorilla")
		wsType = Gorilla
	}

	return newServer(wsType, url, port)
}

func newServer(wsType, url, port string) (websockettypes.Server, error) {
	switch wsType {
	case Gorilla:
		return gorilla.NewServer(url, port)
	default:
		return nil, ErrUnknownWebsocketType{wsType: wsType}
	}
}
