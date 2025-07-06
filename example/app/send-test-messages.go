package main

import (
	"app/actions"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func sendTestMessages(snsPublisher contracts.Publisher) {
	frameworkAction.PublishFromStruct[actions.BlogReceivedAction](snsPublisher, "new-posts", "blog-received", actions.BlogReceivedAction{
		CreatedAt: time.Now(),
		CreatedBy: "John Doe",
		Blog:      "This is the blog post",
	}, nil)
}
