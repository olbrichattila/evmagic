package actionHandlers

import (
	"app/actions"
	"database/sql"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogFailedProfanityHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedSpamCheckAction](message)
	fmt.Println(act.AsAction().FailedAt, act.AsAction().Reason, err)

	return nil, nil
}
