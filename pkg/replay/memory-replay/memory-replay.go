package memoryReplay

import (
	"fmt"
	"sync"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func New(publisher contracts.Publisher) contracts.Replay {
	return &replay{
		publisher:  publisher,
		eventStore: make(map[string][][]byte, 0),
	}
}

type replay struct {
	mu         sync.Mutex
	eventStore map[string][][]byte
	publisher  contracts.Publisher
}

// Replay event if it is a duplicate
func (r *replay) Replay(actionId string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if eventData, ok := r.eventStore[actionId]; ok {
		for _, data := range eventData {

			aInfo, err := frameworkAction.ActionInfoFromPayload(data)
			if err != nil {
				return false, err
			}
			fmt.Println("Duplicate event found", aInfo.MessageIdentifier)
			err = r.publisher.Publish(aInfo.Topic, data)
			if err != nil {
				return false, err
			}
		}

		return true, nil
	}

	return false, nil
}

// Store event after success if need replaying later
func (r *replay) Store(parentActionIdentifier string, replayData []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.eventStore[parentActionIdentifier] = append(r.eventStore[parentActionIdentifier], replayData)
	return nil
}
