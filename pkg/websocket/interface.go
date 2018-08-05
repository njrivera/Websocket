package websocket

import (
	"log"
	"os"

	"github.com/Websocket/internal/gorilla"
	"github.com/Websocket/pkg/websockettypes"
)

func NewWsServer(url string, port string) (websockettypes.Server, <-chan int) {
	wsType := os.Getenv("WS_TYPE")
	if wsType == "" {
		log.Printf("WS_TYPE not set - defaulting to gorilla")
	}
	switch wsType {
	case "gorilla":
		return gorilla.NewWsServer(url, port)
	default:
		return gorilla.NewWsServer(url, port)
	}
}
