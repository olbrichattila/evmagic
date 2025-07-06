package actionHandlers

import (
	"app/actions"
	"database/sql"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogFailedSpamHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedSpamCheckAction](message)
	fmt.Println(act.AsAction().FailedAt, act.AsAction().Reason, err)

	newAct, err := frameworkAction.New[actions.FailedSpamCheckAction](
		"t", "at", actions.FailedSpamCheckAction{}, act.ActionData(),
	)

	bts, _ := newAct.AsBytes()
	return []contracts.ActionResult{
		{Topic: newAct.Topic(), Body: bts},
	}, nil

}
