package sqs

import (
	"context"
	"net/url"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/aws/aws-sdk-go-v2/aws"
	amazonsqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	"github.com/samber/lo"
)

// TODO get those from env
const (
	awsURI       = "http://localhost:4566"
	awsRegion    = "us-east-1"
	awsAccountID = "000000000000"
)

func New() (contracts.Handler, error) {
	h := &handler{}
	err := h.setHandlerConf()
	if err != nil {
		return nil, err
	}

	return h, nil
}

type handler struct {
	subscriber message.Subscriber
	router     *message.Router
}

// Run implements contracts.Handler.
func (h *handler) Run() error {
	return h.router.Run(context.Background())
}

// Handle implements contracts.Handler.
func (h *handler) Handle(topic string, hf contracts.HandlerFunc) {
	// shell I work with the returned *message.handler?
	h.router.AddHandler(
		topic,
		topic, // SQS queue to subscribe to
		h.subscriber,
		topic, // SQS queue to publish to
		nil,
		func(msg *message.Message) ([]*message.Message, error) {
			hf(topic, msg.Payload)
			return nil, nil
		},
	)
}

func (h *handler) setHandlerConf() error {
	logger := watermill.NewStdLogger(false, false)

	sqsOpts := []func(*amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse(awsURI)),
			},
		}),
	}

	subscriberConfig := sqs.SubscriberConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns: sqsOpts,
	}

	var err error
	h.subscriber, err = sqs.NewSubscriber(subscriberConfig, logger)
	if err != nil {
		return err
	}

	h.router, err = message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	return nil
}
