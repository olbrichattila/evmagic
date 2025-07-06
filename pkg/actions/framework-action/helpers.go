package frameworkAction

import (
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	"github.com/olbrichattila/evmagic/pkg/helpers"
)

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

func CreateActionResult[T any](topic, actionType string, action any) contracts.ActionResult {
	// TODO error handling
	failedAction, _ := New[T](topic, actionType, action.(T))
	bytesRes, _ := failedAction.AsBytes()

	return contracts.ActionResult{
		Topic: failedAction.Topic(),
		Body:  bytesRes,
	}
}

func PublishFromStruct[T any](publisher contracts.Publisher, topic, actionType string, action any) {
	act, _ := New[T](topic, actionType, action)
	Publish[T](publisher, act)
}

func Publish[T any](publisher contracts.Publisher, action any) {
	act := action.(Action[T])
	actionBytes, _ := act.AsBytes()
	publisher.Publish(act.Topic(), actionBytes)
}
