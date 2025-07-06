package actionHandlers

import (
	"app/actions"
	"database/sql"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogReceivedSpamHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	fmt.Println(act.AsAction().Blog, act.AsAction().CreatedAt, err)

	_, err = tx.Exec("INSERT INTO blogs (created_at, created_by, blog) VALUES (?, ?, ?)", act.AsAction().CreatedAt, act.AsAction().CreatedBy, act.AsAction().Blog)
	fmt.Println(err)

	return createFailedActionResult[actions.BlogReceivedAction]("spam-failed", "Spam test reason", act.ActionData()), nil

}
