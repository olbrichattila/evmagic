package sqs

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/aws/aws-sdk-go-v2/aws"
	amazonsqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	baseHandler "github.com/olbrichattila/evmagic/pkg/connector/handler"

	"github.com/samber/lo"
)

// TODO get those from env
const (
	awsURI       = "http://localhost:4566"
	awsRegion    = "us-east-1"
	awsAccountID = "000000000000"
)

func New(replay contracts.Replay, db *sql.DB) (contracts.Handler, error) {
	h := &handler{
		replay: replay,
		Handler: baseHandler.Handler{
			Db:     db,
			Topics: map[string]map[string]contracts.HandlerFunc{},
		},
	}
	err := h.setHandlerConf()
	if err != nil {
		return nil, err
	}

	return h, nil
}

type handler struct {
	baseHandler.Handler
	replay     contracts.Replay
	subscriber message.Subscriber
	router     *message.Router
}

// Run implements contracts.Handler.
func (h *handler) Run() error {
	return h.router.Run(context.Background())
}

func (h *handler) Handlers(hd ...contracts.HandlerDef) {
	for _, hDef := range hd {
		h.Handle(hDef.Topic, hDef.ActionType, hDef.Publisher, hDef.HandlerFunc)
	}
}

func (h *handler) Handle(topic, actionType string, publisher contracts.Publisher, hf contracts.HandlerFunc) {
	h.Handler.InternalHandle(h.router, h.replay, h.subscriber, publisher, topic, actionType, hf)
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
