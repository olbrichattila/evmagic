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
}

type ActionBase[T any] struct {
	CorrelationId     string `json:"correlationId"`
	CausationId       string `json:"causationId"`
	MessageIdentifier string `json:"messageIdentifier"`
	ActionType        string `json:"actionType"`
	Content           T      `json:"content"`
}

type SnsAction struct {
	TopicArn string
	Subject  string
	Message  string
}

type actionType struct {
	ActionType string `json:"actionType"`
}

type action[T any] struct {
	base ActionBase[T]
}

func AsSNSActionFromPayload(data []byte) (SnsAction, error) {
	return helpers.ToStruct[SnsAction](data)
}

func ActionTypeFromPayload(data []byte) (string, error) {
	act, err := AsSNSActionFromPayload(data)
	if err != nil {
		return "", err
	}

	res, err := helpers.ToStruct[actionType]([]byte(act.Message))
	if err != nil {
		return "", err
	}

	return res.ActionType, nil

}

func New[T any](actionType string, content any) (Action[T], error) {
	result := action[T]{}
	result.base = ActionBase[T]{
		CorrelationId:     uuid.NewString(),
		CausationId:       uuid.NewString(),
		MessageIdentifier: uuid.NewString(),
		ActionType:        actionType,
		Content:           content.(T),
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
