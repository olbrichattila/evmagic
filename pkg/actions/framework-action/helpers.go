package frameworkAction

import (
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	"github.com/olbrichattila/evmagic/pkg/helpers"
)

func AsSNSActionFromPayload(data []byte) (SnsAction, error) {
	return helpers.ToStruct[SnsAction](data)
}

func ActionInfoFromSNSPayload(data []byte) (*ActionData, error) {
	act, err := AsSNSActionFromPayload(data)
	if err != nil {
		return nil, err
	}

	res, err := helpers.ToStruct[ActionData]([]byte(act.Message))
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func ActionInfoFromPayload(data []byte) (*ActionData, error) {
	res, err := helpers.ToStruct[ActionData](data)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreateActionResult[T any](topic, actionType string, action any, parent *ActionData) contracts.ActionResult {
	// TODO error handling
	failedAction, _ := New[T](topic, actionType, action.(T), parent)
	bytesRes, _ := failedAction.AsBytes()

	return contracts.ActionResult{
		Topic: failedAction.Topic(),
		Body:  bytesRes,
	}
}

func PublishFromStruct[T any](publisher contracts.Publisher, topic, actionType string, action any, parent *ActionData) {
	act, _ := New[T](topic, actionType, action, parent)
	Publish[T](publisher, act)
}

// Publish (Not safe by itself as without parent it can break the correlation tree, use PublishFromStruct with parent instead if possible)
func Publish[T any](publisher contracts.Publisher, action any) {
	act := action.(Action[T])
	actionBytes, _ := act.AsBytes()
	publisher.Publish(act.Topic(), actionBytes)
}
