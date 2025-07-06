package actionHandlers

import (
	"app/actions"
	"database/sql"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogReceivedSpamHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogCheckAction](message)
	// fmt.Println(act.AsAction().BlogID, err)

	_, err = tx.Exec("UPDATE blogs SET banned = 1 WHERE id = ?", act.AsAction().BlogID)
	if err != nil {
		return nil, err
	}

	return createFailedActionResult(act.AsAction().BlogID, "spam-failed", "Spam test reason", "Spam", act.ActionData()), nil
}
