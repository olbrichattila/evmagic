package actionHandlers

import (
	"app/actions"
	"app/entities"
	"database/sql"
	"fmt"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	dbHelper "github.com/olbrichattila/evmagic/pkg/database/dbhelper"
	"github.com/olbrichattila/evmagic/pkg/entity"
)

func BlogReceivedProfanityHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogCheckAction](message)

	dbH := dbHelper.New(tx)
	blog, err := entity.ById[entities.Blogs](dbH, act.AsAction().BlogID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	blog.Banned = true
	err = entity.Save(dbH, blog)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return createFailedActionResult(act.AsAction().BlogID, "profanity-failed", "Profanity test reason", "Profanity", act.ActionData()), nil
}
