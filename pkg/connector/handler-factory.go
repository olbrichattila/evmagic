package connector

import (
	"database/sql"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

type handlers struct {
	registered map[QueueType]contracts.Handler
}

var handlerCache *handlers

func Handler(qt QueueType, rp contracts.Replay, db *sql.DB) (contracts.Handler, error) {
	if handlerCache == nil {
		handlerCache = &handlers{
			registered: make(map[QueueType]contracts.Handler, 0),
		}
	}

	if h, ok := handlerCache.registered[qt]; ok {
		return h, nil
	}

	hb, err := getHandler(qt, rp, db)
	if err != nil {
		return nil, err
	}

	handlerCache.registered[qt] = hb

	return hb, nil
}
