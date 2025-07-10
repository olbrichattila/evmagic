package main

import (
	"database/sql"

	"github.com/olbrichattila/evmagic/pkg/connector"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	"github.com/olbrichattila/evmagic/pkg/replay"
)

func initQueues(db *sql.DB) (contracts.Handler, contracts.Publisher, error) {
	snsPublisher, err := connector.Publisher(connector.TypeSNS)
	if err != nil {
		return nil, nil, err
	}

	replay := replay.New(snsPublisher, db)
	hr := connector.New()
	sqsHandler, err := hr.Handler(connector.TypeSQS, replay, db)
	if err != nil {
		return nil, nil, err
	}

	return sqsHandler, snsPublisher, nil
}
