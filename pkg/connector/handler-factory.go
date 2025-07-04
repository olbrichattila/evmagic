package connector

import (
	"encoding/json"

	actionHelper "github.com/olbrichattila/evmagic/pkg/actions/action-helpers"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

type handlers struct {
	registered map[QueueType]contracts.Handler
}

var handlerCache *handlers

func Handler(qt QueueType) (contracts.Handler, error) {
	if handlerCache == nil {
		handlerCache = &handlers{
			registered: make(map[QueueType]contracts.Handler, 0),
		}
	}

	if h, ok := handlerCache.registered[qt]; ok {
		return h, nil
	}

	hb, err := getHandler(qt)
	if err != nil {
		return nil, err
	}

	handlerCache.registered[qt] = hb

	return hb, nil
}

func AsSNSAction(data []byte) (actionHelper.SnsAction, error) {
	return actionHelper.ToSNSAction(data)
}

func AsAction[T any](data []byte) (T, error) {
	var res T
	act, err := actionHelper.ToSNSAction(data)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(act.Message), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
