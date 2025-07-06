package actionHandlers

import (
	"app/actions"
	"database/sql"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogReceivedPlagiarismHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	fmt.Println(act.AsAction().Blog, act.AsAction().CreatedAt, err)

	r := createFailedActionResult[actions.BlogReceivedAction]("plagiarism-failed", "Plagiarism test reason", act.ActionData())

	// Test replaying the same event
	// bts, _ := act.AsBytes()
	// r = append(r, contracts.ActionResult{Topic: act.Topic(), Body: bts})
	return r, nil
}
