package actionHandlers

import (
	"app/actions"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
)

func BlogReceivedHandler(message []byte) [][]byte {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	fmt.Println(act.AsAction().Blog, act.AsAction().CreatedAt, err)

	return nil
}
