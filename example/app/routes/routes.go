package routes

import (
	actionHandlers "app/action-handlers"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func InitRoutes(handler contracts.Handler, publisher contracts.Publisher) {
	handler.Handlers(
		contracts.HandlerDef{
			Topic:       "check-failed-queue",
			ActionType:  "blog-received",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogReceivedHandler,
		},
		contracts.HandlerDef{
			Topic:       "spam-check-queue",
			ActionType:  "blog-received",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogReceivedSpamHandler,
		},
		contracts.HandlerDef{
			Topic:       "profanity-check-queue",
			ActionType:  "blog-received",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogReceivedProfanityHandler,
		},
		contracts.HandlerDef{
			Topic:       "plagiarism-check-queue",
			ActionType:  "blog-received",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogReceivedPlagiarismHandler,
		},
		// failed
		contracts.HandlerDef{
			Topic:       "check-failed-queue",
			ActionType:  "spam-failed",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogFailedSpamHandler,
		},
		contracts.HandlerDef{
			Topic:       "check-failed-queue",
			ActionType:  "profanity-failed",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogFailedProfanityHandler,
		},
		contracts.HandlerDef{
			Topic:       "check-failed-queue",
			ActionType:  "plagiarism-failed",
			Publisher:   publisher,
			HandlerFunc: actionHandlers.BlogFailedPlagiarismHandler,
		},
	)
}
