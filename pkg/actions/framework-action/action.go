// Package frameworkAction is a helper function to create an action from a struct which can be published to the queue
// including Correlation ID, Causation ID, and Action type for internal routing
// TODO add a function called Parent, if it is set, adjust the Correlation ID
package frameworkAction

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/olbrichattila/evmagic/pkg/helpers"
)

type Action[T any] interface {
	AsBytes() ([]byte, error)
	AsAction() T
	CorrelationId() string
	CausationId() string
	MessageIdentifier() string
	ActionType() string
	Topic() string
	ActionData() *ActionData
}

type SnsAction struct {
	TopicArn string
	Subject  string
	Message  string
}

type ActionData struct {
	Topic             string `json:"topic"`
	CorrelationId     string `json:"correlationId"`
	CausationId       string `json:"causationId"`
	MessageIdentifier string `json:"messageIdentifier"`
	ActionType        string `json:"actionType"`
}

type ActionBase[T any] struct {
	ActionData
	Content T `json:"content"`
}

type action[T any] struct {
	base ActionBase[T]
}

func New[T any](topic, actionType string, content any, parentActionData *ActionData) (Action[T], error) {
	result := action[T]{}
	result.base = ActionBase[T]{
		ActionData: ActionData{
			Topic:             topic,
			MessageIdentifier: uuid.NewString(),
			CorrelationId:     "",
			CausationId:       "",

			ActionType: actionType,
		},
		Content: content.(T),
	}

	if parentActionData != nil {
		result.base.CausationId = parentActionData.MessageIdentifier
		result.base.CorrelationId = parentActionData.CorrelationId
	} else {
		result.base.CausationId = result.base.MessageIdentifier
		result.base.CorrelationId = result.base.MessageIdentifier
	}

	return result, nil
}

// NewFromPayload implements Action.
func NewFromPayload[T any](payload []byte) (Action[T], error) {
	act, err := AsSNSActionFromPayload(payload)
	if err != nil {
		return nil, err
	}

	res, err := helpers.ToStruct[ActionBase[T]]([]byte(act.Message))
	if err != nil {
		return nil, err
	}

	return action[T]{
		base: res,
	}, nil
}

// AsBytes implements Action.
func (a action[T]) AsBytes() ([]byte, error) {
	return json.Marshal(a.base)
}

func (a action[T]) AsAction() T {
	return a.base.Content
}

func (a action[T]) CorrelationId() string {
	return a.base.CorrelationId
}

func (a action[T]) CausationId() string {
	return a.base.CausationId
}

func (a action[T]) MessageIdentifier() string {
	return a.base.MessageIdentifier
}

func (a action[T]) ActionType() string {
	return a.base.ActionType
}

func (a action[T]) Topic() string {
	return a.base.Topic
}

func (a action[T]) ActionData() *ActionData {
	return &a.base.ActionData
}
