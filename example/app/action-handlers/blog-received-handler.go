package actionHandlers

import (
	"app/actions"
	"database/sql"
	"log"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func BlogReceivedHandler(tx *sql.Tx, message []byte) ([]contracts.ActionResult, error) {
	act, err := frameworkAction.NewFromPayload[actions.BlogReceivedAction](message)
	// fmt.Println(act.AsAction().Blog, act.AsAction().CreatedAt, err)

	//	fmt.Println("Store blog!!")
	result, err := tx.Exec("INSERT INTO blogs (created_at, created_by, blog) VALUES (?, ?, ?)", act.AsAction().CreatedAt, act.AsAction().CreatedBy, act.AsAction().Blog)
	//	fmt.Println(err)

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal("LastInsertId error:", err)
		return nil, err
	}

	return []contracts.ActionResult{
		frameworkAction.CreateActionResult[actions.BlogCheckAction]("new-posts", "blog-received", actions.BlogCheckAction{
			CreatedAt: time.Now(),
			BlogID:    lastID,
		}, act.ActionData()),
	}, nil
}
