package types

type Server interface {
	Broadcast(msg interface{})
	ListenForNewClients() <-chan Client
}
