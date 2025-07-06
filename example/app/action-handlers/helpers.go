package actionHandlers

import (
	"app/actions"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func createFailedActionResult[T any](actionType, reason string, parentAct *frameworkAction.ActionData) []contracts.ActionResult {
	return []contracts.ActionResult{
		frameworkAction.CreateActionResult[actions.FailedPlagiarismCheckAction]("post-processed", actionType, actions.FailedPlagiarismCheckAction{
			FailedAt: time.Now(),
			Reason:   reason,
		}, parentAct),
	}
}
