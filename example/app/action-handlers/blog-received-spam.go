package actionHandlers

import (
	"app/actions"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogReceivedSpamHandler(message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	fmt.Println(act.AsAction().Blog, act.AsAction().CreatedAt, err)

	return createFailedActionResult("spam-failed", "Spam test reason"), nil

}
