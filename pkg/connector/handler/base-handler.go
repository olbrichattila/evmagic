package baseHandler

import (
	"database/sql"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

type Handler struct {
	Db     *sql.DB
	Topics map[string]map[string]contracts.HandlerFunc
}

// InternalHandle implements contracts.Handler.
func (h *Handler) InternalHandle(router *message.Router, replay contracts.Replay, subscriber message.Subscriber, publisher contracts.Publisher, topic, actionType string, hf contracts.HandlerFunc) {
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
			tx, err := h.Db.Begin()
			if err != nil {
				// TODO log
				fmt.Println("Error creating transaction")
				return nil, fmt.Errorf("invalid action type: %w", err)
			}

			actionInfo, err := frameworkAction.ActionInfoFromSNSPayload(msg.Payload)
			if err != nil {
				// TODO log
				tx.Rollback()
				return nil, fmt.Errorf("invalid action type: %w", err)
			}

			// Message replay
			replayed, err := replay.Replay(topic + "_" + actionInfo.MessageIdentifier)
			if err != nil {
				// TODO log
				tx.Rollback()
				return nil, fmt.Errorf("invalid action type: %w", err)
			}

			if replayed {
				tx.Rollback()
				return nil, nil
			}

			if handleFnc, ok := h.Topics[topic][actionInfo.ActionType]; ok {
				actionsToPublish, err := handleFnc(tx, msg.Payload)
				if err != nil {
					// TODO log
					tx.Rollback()
					return nil, err
				}

				err = replay.Register(topic + "_" + actionInfo.MessageIdentifier)
				if err != nil {
					// TODO log
					tx.Rollback()
					return nil, err
				}

				tx.Commit()

				// Publish child actions
				for _, msgToPub := range actionsToPublish {
					err := publisher.Publish(msgToPub.Topic, msgToPub.Body)
					if err != nil {
						// TODO log
						return nil, err
					}

					err = replay.Store(topic+"_"+actionInfo.MessageIdentifier, msgToPub.Body)
					if err != nil {
						// TODO log
						return nil, err
					}
				}

				return nil, nil
			}

			// TODO log
			return nil, fmt.Errorf("action type does not have handler: %s", actionType)
		},
	)
}
