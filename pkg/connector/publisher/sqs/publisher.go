package sqs

import (
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

func New() (contracts.Publisher, error) {
	pb := &publisher{}

	if err := pb.setPublisher(); err != nil {
		return nil, err
	}

	return pb, nil
}

type publisher struct {
	publisher message.Publisher
}

// Publish implements contracts.Publisher.
func (p *publisher) Publish(topic string, msg []byte) error {
	newMsg := message.NewMessage(watermill.NewUUID(), msg)

	return p.publisher.Publish(topic, newMsg)
}

func (p *publisher) setPublisher() error {
	logger := watermill.NewStdLogger(false, false)

	sqsOpts := []func(*amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse(awsURI)),
			},
		}),
	}

	publisherConfig := sqs.PublisherConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns: sqsOpts,
	}

	var err error
	p.publisher, err = sqs.NewPublisher(publisherConfig, logger)
	if err != nil {
		return err
	}

	return nil
}
