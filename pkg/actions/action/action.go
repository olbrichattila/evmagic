package action

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ActionBase struct {
	CorrelationId     string `json:"correlationId"`
	CausationId       string `json:"causationId"`
	MessageIdentifier string `json:"messageIdentifier"`
	ActionType        string `json:"actionType"`
}

type ActionContent struct {
	ActionBase
	Content any `json:"content"`
}

func New(actionType string, content any) ([]byte, error) {
	act := ActionContent{
		ActionBase: ActionBase{
			CorrelationId:     uuid.NewString(),
			CausationId:       uuid.NewString(),
			MessageIdentifier: uuid.NewString(),
			ActionType:        actionType,
		},
		Content: content,
	}

	return json.Marshal(act)
}
