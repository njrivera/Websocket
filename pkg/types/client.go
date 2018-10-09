package types

type Client interface {
	Send(msg interface{}) error
	Receive() ([]byte, error)
	Close() error
}
