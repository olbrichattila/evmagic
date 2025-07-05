package main

import (
	"app/actions"
	"time"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func sendTestMessages(snsPublisher contracts.Publisher) {
	helloAct, _ := frameworkAction.New[actions.BlogReceivedAction]("blog-received", actions.BlogReceivedAction{
		CreatedAt: time.Now(),
		CreatedBy: "John Doe",
		Blog:      "This is the blog post",
	})
	actionBytes, _ := helloAct.AsBytes()

	helloAct2, _ := frameworkAction.New[actions.BlogReceivedAction]("blog-received-2", actions.BlogReceivedAction{
		CreatedAt: time.Now(),
		CreatedBy: "John Doe2",
		Blog:      "This is the blog post2",
	})
	actionBytes2, _ := helloAct2.AsBytes()

	snsPublisher.Publish("new-posts", actionBytes)
	snsPublisher.Publish("new-posts", actionBytes)
	snsPublisher.Publish("new-posts", actionBytes2)

}
