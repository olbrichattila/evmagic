package actionHandlers

import (
	"app/actions"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogFailedSpamHandler(message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedSpamCheckAction](message)
	fmt.Println(act.AsAction().FailedAt, act.AsAction().Reason, err)

	return nil, nil
}
