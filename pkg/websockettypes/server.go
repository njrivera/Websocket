package websockettypes

type Server interface {
	Send(id int, msg interface{}) error
	Receive(id int) ([]byte, error)
	Close(id int) error
}
