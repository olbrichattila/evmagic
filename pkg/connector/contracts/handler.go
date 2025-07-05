package contracts

type HandlerFunc func(msg []byte) [][]byte

type HandlerDef struct {
	Topic       string
	ActionType  string
	HandlerFunc HandlerFunc
}

type Handler interface {
	Handle(topic, actionType string, hf HandlerFunc)
	Handlers(hd ...HandlerDef)
	Run() error
}
