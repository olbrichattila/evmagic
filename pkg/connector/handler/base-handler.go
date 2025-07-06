package baseHandler

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

type Handler struct {
	Topics map[string]map[string]contracts.HandlerFunc
}

// InternalHandle implements contracts.Handler.
func (h *Handler) InternalHandle(router *message.Router, subscriber message.Subscriber, publisher contracts.Publisher, topic, actionType string, hf contracts.HandlerFunc) {
	if existingTopic, ok := h.Topics[topic]; ok {
		existingTopic[actionType] = hf
		return
	}

	h.Topics[topic] = map[string]contracts.HandlerFunc{
		actionType: hf,
	}

	router.AddHandler(
		topic,
		topic, // SQS queue to subscribe to
		subscriber,
		topic, // SQS queue to publish to
		nil,
		func(msg *message.Message) ([]*message.Message, error) {
			actionType, err := frameworkAction.ActionTypeFromPayload(msg.Payload)
			if err != nil {
				// TODO log
				return nil, fmt.Errorf("invalid action type: %w", err)
			}

			if handleFnc, ok := h.Topics[topic][actionType]; ok {
				actionsToPublish, err := handleFnc(msg.Payload)
				if err == nil {
					for _, msgToPub := range actionsToPublish {
						err := publisher.Publish(msgToPub.Topic, msgToPub.Body)
						if err != nil {
							return nil, err
						}
					}
				}
				return nil, err
			}

			// TODO log
			return nil, fmt.Errorf("action type does not have handler: %s", actionType)
		},
	)
}
