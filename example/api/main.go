package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sns"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/samber/lo"

	amazonsns "github.com/aws/aws-sdk-go-v2/service/sns"
	transport "github.com/aws/smithy-go/endpoints"
)

const (
	awsURI       = "http://localhost:4566"
	awsRegion    = "us-east-1"
	awsAccountID = "000000000000"
	topicName    = "new-posts"
)

type blogAction struct {
	Blog string `json:"blog"`
}

func main() {
	sendMessage(getPrompt())
}

func getPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := ""

	for {
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			break
		}

		lines = lines + text

	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading input: %s", err.Error()))
	}

	return lines
}

func sendMessage(messageToPublish string) {
	action := blogAction{Blog: messageToPublish}
	messageJSON, _ := json.Marshal(action)
	publisher := getSNSPublisher()

	msg := message.NewMessage(watermill.NewUUID(), messageJSON)
	if err := publisher.Publish(topicName, msg); err != nil {
		panic(err)
	}
}

func getSNSPublisher() message.Publisher {
	logger := watermill.NewStdLogger(false, false)
	snsOpts := []func(*amazonsns.Options){
		amazonsns.WithEndpointResolverV2(sns.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse(awsURI)),
			},
		}),
	}

	topicResolver, err := sns.NewGenerateArnTopicResolver(awsAccountID, awsRegion)
	if err != nil {
		panic(err)
	}

	publisherConfig := sns.PublisherConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns:        snsOpts,
		TopicResolver: topicResolver,
	}

	publisher, err := sns.NewPublisher(publisherConfig, logger)
	if err != nil {
		panic(err)
	}

	return publisher
}
