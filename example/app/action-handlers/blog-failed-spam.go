package actionHandlers

import (
	"app/actions"
	"app/entities"
	"database/sql"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	dbHelper "github.com/olbrichattila/evmagic/pkg/database/dbhelper"
	"github.com/olbrichattila/evmagic/pkg/entity"
)

func BlogFailedSpamHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedCheckAction](message)
	if err != nil {
		return nil, err
	}

	err = entity.Save(dbHelper.New(tx), entities.BlogCheck{
		BlogId:    act.AsAction().BlogID,
		CheckType: act.AsAction().CheckType,
		Reason:    act.AsAction().Reason,
		CreatedAt: act.AsAction().FailedAt,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil

}
