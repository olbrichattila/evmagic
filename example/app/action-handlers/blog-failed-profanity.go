package actionHandlers

import (
	"app/actions"
	"database/sql"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogFailedProfanityHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedCheckAction](message)
	// fmt.Println(act.AsAction().FailedAt, act.AsAction().Reason, act.AsAction().CheckType, err)

	_, err = tx.Exec("INSERT INTO blog_checks (blog_id, check_type, reason, created_at) VALUES (?, ?, ?, ?)", act.AsAction().BlogID, act.AsAction().CheckType, act.AsAction().Reason, act.AsAction().FailedAt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
