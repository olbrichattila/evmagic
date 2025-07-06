package contracts

import "database/sql"

type ActionResult struct {
	Topic string
	Body  []byte
}
type HandlerFunc func(tx *sql.Tx, msg []byte) ([]ActionResult, error)

type HandlerDef struct {
	Topic       string
	ActionType  string
	Publisher   Publisher
	HandlerFunc HandlerFunc
}

type Handler interface {
	Handle(topic, actionType string, publisher Publisher, hf HandlerFunc)
	Handlers(hd ...HandlerDef)
	Run() error
}
