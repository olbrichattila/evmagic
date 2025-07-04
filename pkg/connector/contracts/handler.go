package contracts

type HandlerFunc func(topic string, msg []byte) [][]byte

type Handler interface {
	Handle(topic string, hf HandlerFunc)
	Run() error
}
