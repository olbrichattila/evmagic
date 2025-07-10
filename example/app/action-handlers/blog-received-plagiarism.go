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

func BlogReceivedPlagiarismHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogCheckAction](message)

	dbH := dbHelper.New(tx)
	blog, err := entity.ById[entities.Blogs](dbH, act.AsAction().BlogID)
	if err != nil {
		fmt.Println("received", err)
		return nil, err
	}

	blog.Banned = true
	err = entity.Save(dbH, blog)
	if err != nil {
		fmt.Println("Update banned", err)
		return nil, err
	}
	return createFailedActionResult(act.AsAction().BlogID, "plagiarism-failed", "Plagiarism test reason", "Plagiarism", act.ActionData()), nil
}
