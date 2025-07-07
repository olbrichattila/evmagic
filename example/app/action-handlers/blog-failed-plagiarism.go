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

func BlogFailedPlagiarismHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.FailedCheckAction](message)
	// fmt.Println(act.AsAction().FailedAt, act.AsAction().Reason, act.AsAction().CheckType, err)

	/// TEST DB helper
	dbH := dbHelper.New(tx)
	res := dbH.QueryAll("select * from blogs")
	for row := range res {
		fmt.Println(row)
	}

	// Test entities
	// dbH := dbHelper.New(tx)
	blogs, err := entity.All[entities.Blogs](dbH)
	fmt.Println(blogs, err)

	_, err = tx.Exec("INSERT INTO blog_checks (blog_id, check_type, reason, created_at) VALUES (?, ?, ?, ?)", act.AsAction().BlogID, act.AsAction().CheckType, act.AsAction().Reason, act.AsAction().FailedAt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
