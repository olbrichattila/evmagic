package actionHandlers

import (
	"app/actions"
	"app/entities"
	"database/sql"
	"fmt"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	dbHelper "github.com/olbrichattila/evmagic/pkg/database/dbhelper"
	"github.com/olbrichattila/evmagic/pkg/entity"
)

func BlogReceivedHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	if err != nil {
		return nil, err
	}

	blogEntity := &entities.Blogs{
		CreatedAt: act.AsAction().CreatedAt.Format(time.DateTime),
		CreatedBy: act.AsAction().CreatedBy,
		Blog:      act.AsAction().Blog,
	}

	err = entity.Save(dbHelper.New(tx), blogEntity)
	if err != nil {
		fmt.Println("???", err)
		return nil, err
	}

	return []contracts.ActionResult{
		frameworkAction.CreateActionResult[actions.BlogCheckAction]("new-posts", "blog-received", actions.BlogCheckAction{
			CreatedAt: time.Now(),
			BlogID:    blogEntity.Id,
		}, act.ActionData()),
	}, nil
}
