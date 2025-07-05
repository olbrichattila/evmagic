package routes

import (
	actionHandlers "app/action-handlers"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func InitRoutes(handler contracts.Handler) {
	handler.Handlers(
		contracts.HandlerDef{Topic: "spam-check-queue", ActionType: "blog-received-2", HandlerFunc: actionHandlers.BlogReceivedHandler},
		contracts.HandlerDef{Topic: "spam-check-queue", ActionType: "blog-received", HandlerFunc: actionHandlers.BlogReceivedHandler},
	)
}
