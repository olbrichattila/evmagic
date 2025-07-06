package main

import (
	"fmt"
	"os"

	"github.com/olbrichattila/evmagic/pkg/connector"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	"github.com/olbrichattila/evmagic/pkg/database/connection"
	"github.com/olbrichattila/evmagic/pkg/replay"
)

func initQueues() (contracts.Handler, contracts.Publisher, error) {
	snsPublisher, err := connector.Publisher(connector.TypeSNS)
	if err != nil {
		return nil, nil, err
	}

	db, err := connection.Open()
	if err != nil {
		fmt.Println("Cannot connect to the database", err.Error())
		os.Exit(1)
	}

	replay := replay.New(snsPublisher, db)
	sqsHandler, err := connector.Handler(connector.TypeSQS, replay, db)
	if err != nil {
		return nil, nil, err
	}

	return sqsHandler, snsPublisher, nil
}
