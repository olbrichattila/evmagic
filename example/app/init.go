package main

import (
	"github.com/olbrichattila/evmagic/pkg/connector"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func initQueues() (contracts.Handler, contracts.Publisher, error) {
	sqsHandler, err := connector.Handler(connector.TypeSQS)
	if err != nil {
		return nil, nil, err
	}

	snsPublisher, err := connector.Publisher(connector.TypeSNS)
	if err != nil {
		return nil, nil, err
	}

	return sqsHandler, snsPublisher, nil
}
