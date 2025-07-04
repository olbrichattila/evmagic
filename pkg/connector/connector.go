// Package connector abstracts connection to different queues
// like ASW SNS, SQS, Kafka, RabbitMq....
package connector

import (
	frameworkErrors "github.com/olbrichattila/evmagic/pkg/connector/framework-errors"
	snsHandler "github.com/olbrichattila/evmagic/pkg/connector/handler/sns"
	sqsHandler "github.com/olbrichattila/evmagic/pkg/connector/handler/sqs"
	snsPublisher "github.com/olbrichattila/evmagic/pkg/connector/publisher/sns"
	sqsPublisher "github.com/olbrichattila/evmagic/pkg/connector/publisher/sqs"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

// QueueType is for types
type QueueType int

// List of queue types, line sns,sqs, Add .... kafka, rabbitMq....
const (
	TypeSNS QueueType = iota
	TypeSQS
)

func getPublisher(qt QueueType) (contracts.Publisher, error) {
	switch qt {
	case TypeSNS:
		return snsPublisher.New()
	case TypeSQS:
		return sqsPublisher.New()
	default:
		return nil, frameworkErrors.ErrInvalidQueueType
	}
}

func getHandler(qt QueueType) (contracts.Handler, error) {
	switch qt {
	case TypeSNS:
		return snsHandler.New()
	case TypeSQS:
		return sqsHandler.New()
	default:
		return nil, frameworkErrors.ErrInvalidQueueType
	}
}
