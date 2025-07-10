package connector

import (
	"database/sql"
	"sync"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func New() *Handlers {
	return &Handlers{
		registered: map[QueueType]contracts.Handler{},
	}

}

type Handlers struct {
	mu         sync.Mutex
	registered map[QueueType]contracts.Handler
}

func (h *Handlers) Handler(qt QueueType, rp contracts.Replay, db *sql.DB) (contracts.Handler, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h, ok := h.registered[qt]; ok {
		return h, nil
	}

	hb, err := getHandler(qt, rp, db)
	if err != nil {
		return nil, err
	}

	h.registered[qt] = hb

	return hb, nil
}
