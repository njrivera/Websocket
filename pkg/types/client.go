package types

type Client interface {
	GetName() string
	Send(msg interface{}) error
	Receive() ([]byte, error)
	Close() error
}
