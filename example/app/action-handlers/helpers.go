package actionHandlers

import (
	"app/actions"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func createFailedActionResult(blogID int64, actionType, reason, checkType string, parentAct *frameworkAction.ActionData) []contracts.ActionResult {
	return []contracts.ActionResult{
		frameworkAction.CreateActionResult[actions.FailedCheckAction]("post-processed", actionType, actions.FailedCheckAction{
			BlogID:    blogID,
			CheckType: checkType,
			FailedAt:  time.Now(),
			Reason:    reason,
		}, parentAct),
	}
}
